package db

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/concourse/concourse/atc"
)

//counterfeiter:generate . WorkerFactory
type WorkerFactory interface {
	GetWorker(name string) (Worker, bool, error)
	SaveWorker(atcWorker atc.Worker, ttl time.Duration) (Worker, error)
	HeartbeatWorker(worker atc.Worker, ttl time.Duration) (Worker, error)
	Workers() ([]Worker, error)
	VisibleWorkers([]string) ([]Worker, error)

	FindWorkersForContainerByOwner(ContainerOwner) ([]Worker, error)
	BuildContainersCountPerWorker() (map[string]int, error)
}

type workerFactory struct {
	conn  DbConn
	cache *WorkerCache
}

func NewWorkerFactory(conn DbConn, cache *WorkerCache) WorkerFactory {
	return &workerFactory{
		conn:  conn,
		cache: cache,
	}
}

var workersQuery = psql.Select(`
		w.name,
		w.version,
		w.addr,
		w.state,
		w.baggageclaim_url,
		w.certs_path,
		w.http_proxy_url,
		w.https_proxy_url,
		w.no_proxy,
		w.active_containers,
		w.active_volumes,
		w.resource_types,
		w.platform,
		w.tags,
		t.name,
		w.team_id,
		w.start_time,
		w.expires,
		w.ephemeral
	`).
	From("workers w").
	LeftJoin("teams t ON w.team_id = t.id")

func (f *workerFactory) GetWorker(name string) (Worker, bool, error) {
	workers, err := f.cache.Workers()
	if err != nil {
		return nil, false, err
	}

	for _, worker := range workers {
		if worker.Name() == name {
			return worker, true, nil
		}
	}

	return nil, false, nil
}

func (f *workerFactory) VisibleWorkers(teamNames []string) ([]Worker, error) {
	workers, err := f.cache.Workers()
	if err != nil {
		return nil, err
	}

	isVisible := func(worker Worker) bool {
		if worker.TeamID() == 0 {
			return true
		}

		for _, team := range teamNames {
			if worker.TeamName() == team {
				return true
			}
		}

		return false
	}

	visibleWorkers := []Worker{}
	for _, worker := range workers {
		if isVisible(worker) {
			visibleWorkers = append(visibleWorkers, worker)
		}
	}

	return visibleWorkers, nil
}

func (f *workerFactory) Workers() ([]Worker, error) {
	return f.cache.Workers()
}

func getWorker(conn DbConn, query sq.SelectBuilder) (Worker, bool, error) {
	row := query.
		RunWith(conn).
		QueryRow()

	w := &worker{conn: conn}

	err := scanWorker(w, row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}
		return nil, false, err
	}

	return w, true, nil
}

func getWorkers(conn DbConn, query sq.SelectBuilder) ([]Worker, error) {
	rows, err := query.RunWith(conn).Query()
	if err != nil {
		return nil, err
	}
	defer Close(rows)

	workers := []Worker{}

	for rows.Next() {
		worker := &worker{conn: conn}
		err := scanWorker(worker, rows)
		if err != nil {
			return nil, err
		}

		workers = append(workers, worker)
	}

	return workers, nil
}

func scanWorker(worker *worker, row scannable) error {
	var (
		version       sql.NullString
		addStr        sql.NullString
		state         string
		bcURLStr      sql.NullString
		certsPathStr  sql.NullString
		httpProxyURL  sql.NullString
		httpsProxyURL sql.NullString
		noProxy       sql.NullString
		resourceTypes []byte
		platform      sql.NullString
		tags          []byte
		teamName      sql.NullString
		teamID        sql.NullInt64
		startTime     sql.NullTime
		expiresAt     sql.NullTime
		ephemeral     sql.NullBool
	)

	err := row.Scan(
		&worker.name,
		&version,
		&addStr,
		&state,
		&bcURLStr,
		&certsPathStr,
		&httpProxyURL,
		&httpsProxyURL,
		&noProxy,
		&worker.activeContainers,
		&worker.activeVolumes,
		&resourceTypes,
		&platform,
		&tags,
		&teamName,
		&teamID,
		&startTime,
		&expiresAt,
		&ephemeral,
	)
	if err != nil {
		return err
	}

	if version.Valid {
		worker.version = &version.String
	}

	if addStr.Valid {
		worker.gardenAddr = &addStr.String
	}

	if bcURLStr.Valid {
		worker.baggageclaimURL = &bcURLStr.String
	}

	if certsPathStr.Valid {
		worker.certsPath = &certsPathStr.String
	}

	worker.state = WorkerState(state)
	worker.startTime = startTime.Time
	worker.expiresAt = expiresAt.Time

	if httpProxyURL.Valid {
		worker.httpProxyURL = httpProxyURL.String
	}

	if httpsProxyURL.Valid {
		worker.httpsProxyURL = httpsProxyURL.String
	}

	if noProxy.Valid {
		worker.noProxy = noProxy.String
	}

	if teamName.Valid {
		worker.teamName = teamName.String
	}

	if teamID.Valid {
		worker.teamID = int(teamID.Int64)
	}

	if platform.Valid {
		worker.platform = platform.String
	}

	if ephemeral.Valid {
		worker.ephemeral = ephemeral.Bool
	}

	err = json.Unmarshal(resourceTypes, &worker.resourceTypes)
	if err != nil {
		return err
	}

	return json.Unmarshal(tags, &worker.tags)
}

func (f *workerFactory) HeartbeatWorker(atcWorker atc.Worker, ttl time.Duration) (Worker, error) {
	// In order to be able to calculate the ttl that we return to the caller
	// we must compare time.Now() to the worker.expires column
	// However, workers.expires column is a "timestamp (without timezone)"
	// So we format time.Now() without any timezone information and then
	// parse that using the same layout to strip the timezone information

	tx, err := f.conn.Begin()
	if err != nil {
		return nil, err
	}
	defer Rollback(tx)

	expires := "NULL"
	if ttl != 0 {
		expires = fmt.Sprintf(`NOW() + '%d second'::INTERVAL`, int(ttl.Seconds()))
	}

	cSQL, _, err := sq.Case("state").
		When("'landing'::worker_state", "'landing'::worker_state").
		When("'landed'::worker_state", "'landed'::worker_state").
		When("'retiring'::worker_state", "'retiring'::worker_state").
		Else("'running'::worker_state").
		ToSql()

	if err != nil {
		return nil, err
	}

	_, err = psql.Update("workers").
		Set("expires", sq.Expr(expires)).
		Set("active_containers", atcWorker.ActiveContainers).
		Set("active_volumes", atcWorker.ActiveVolumes).
		Set("state", sq.Expr("("+cSQL+")")).
		Where(sq.Eq{"name": atcWorker.Name}).
		RunWith(tx).
		Exec()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrWorkerNotPresent
		}
		return nil, err
	}

	row := workersQuery.Where(sq.Eq{"w.name": atcWorker.Name}).
		RunWith(tx).
		QueryRow()

	worker := &worker{conn: f.conn}
	err = scanWorker(worker, row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrWorkerNotPresent
		}
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return worker, nil

}

func (f *workerFactory) SaveWorker(atcWorker atc.Worker, ttl time.Duration) (Worker, error) {
	tx, err := f.conn.Begin()
	if err != nil {
		return nil, err
	}

	defer Rollback(tx)

	savedWorker, err := saveWorker(tx, atcWorker, nil, ttl, f.conn)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return savedWorker, nil
}

func (f *workerFactory) FindWorkersForContainerByOwner(owner ContainerOwner) ([]Worker, error) {
	ownerQuery, found, err := owner.Find(f.conn)
	if err != nil {
		return nil, err
	}

	if !found {
		return []Worker{}, nil
	}

	ownerEq := sq.Eq{}
	for k, v := range ownerQuery {
		ownerEq["c."+k] = v
	}

	workers, err := getWorkers(f.conn, workersQuery.Join("containers c ON c.worker_name = w.name").Where(sq.And{
		ownerEq,
	}))
	if err != nil {
		return nil, err
	}

	return workers, nil
}

func (f *workerFactory) BuildContainersCountPerWorker() (map[string]int, error) {
	return f.cache.WorkerContainerCounts()
}

func saveWorker(tx Tx, atcWorker atc.Worker, teamID *int, ttl time.Duration, conn DbConn) (Worker, error) {
	resourceTypes, err := json.Marshal(atcWorker.ResourceTypes)
	if err != nil {
		return nil, err
	}

	tags, err := json.Marshal(atcWorker.Tags)
	if err != nil {
		return nil, err
	}

	expires := "NULL"
	if ttl != 0 {
		expires = fmt.Sprintf(`NOW() + '%d second'::INTERVAL`, int(ttl.Seconds()))
	}

	startTime := fmt.Sprintf(`to_timestamp(%d)`, atcWorker.StartTime)

	var workerState WorkerState
	if atcWorker.State != "" {
		workerState = WorkerState(atcWorker.State)
	} else {
		workerState = WorkerStateRunning
	}

	var workerVersion *string
	if atcWorker.Version != "" {
		workerVersion = &atcWorker.Version
	}

	values := []any{
		atcWorker.GardenAddr,
		atcWorker.ActiveContainers,
		atcWorker.ActiveVolumes,
		resourceTypes,
		tags,
		atcWorker.Platform,
		atcWorker.BaggageclaimURL,
		atcWorker.CertsPath,
		atcWorker.HTTPProxyURL,
		atcWorker.HTTPSProxyURL,
		atcWorker.NoProxy,
		atcWorker.Name,
		workerVersion,
		string(workerState),
		teamID,
		atcWorker.Ephemeral,
	}

	conflictValues := values
	var matchTeamUpsert string
	if teamID == nil {
		matchTeamUpsert = "workers.team_id IS NULL"
	} else {
		matchTeamUpsert = "workers.team_id = ?"
		conflictValues = append(conflictValues, *teamID)
	}

	rows, err := psql.Insert("workers").
		Columns(
			"expires",
			"start_time",
			"addr",
			"active_containers",
			"active_volumes",
			"resource_types",
			"tags",
			"platform",
			"baggageclaim_url",
			"certs_path",
			"http_proxy_url",
			"https_proxy_url",
			"no_proxy",
			"name",
			"version",
			"state",
			"team_id",
			"ephemeral",
		).
		Values(append([]any{
			sq.Expr(expires),
			sq.Expr(startTime),
		}, values...)...).
		Suffix(`
			ON CONFLICT (name) DO UPDATE SET
				expires = `+expires+`,
				start_time = `+startTime+`,
				addr = ?,
				active_containers = ?,
				active_volumes = ?,
				resource_types = ?,
				tags = ?,
				platform = ?,
				baggageclaim_url = ?,
				certs_path = ?,
				http_proxy_url = ?,
				https_proxy_url = ?,
				no_proxy = ?,
				name = ?,
				version = ?,
				state = ?,
				team_id = ?,
				ephemeral = ?
			WHERE `+matchTeamUpsert,
			conflictValues...,
		).
		RunWith(tx).
		Exec()
	if err != nil {
		return nil, err
	}

	count, err := rows.RowsAffected()
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, errors.New("worker already exists and is either global or owned by another team")
	}

	var workerTeamID int
	if teamID != nil {
		workerTeamID = *teamID
	}

	savedWorker := &worker{
		name:             atcWorker.Name,
		version:          workerVersion,
		state:            workerState,
		gardenAddr:       &atcWorker.GardenAddr,
		baggageclaimURL:  &atcWorker.BaggageclaimURL,
		certsPath:        atcWorker.CertsPath,
		httpProxyURL:     atcWorker.HTTPProxyURL,
		httpsProxyURL:    atcWorker.HTTPSProxyURL,
		noProxy:          atcWorker.NoProxy,
		activeContainers: atcWorker.ActiveContainers,
		activeVolumes:    atcWorker.ActiveVolumes,
		resourceTypes:    atcWorker.ResourceTypes,
		platform:         atcWorker.Platform,
		tags:             atcWorker.Tags,
		teamName:         atcWorker.Team,
		teamID:           workerTeamID,
		startTime:        time.Unix(atcWorker.StartTime, 0),
		ephemeral:        atcWorker.Ephemeral,
		conn:             conn,
	}

	workerBaseResourceTypeIDs := []int{}

	for _, resourceType := range atcWorker.ResourceTypes {
		workerResourceType := WorkerResourceType{
			Worker:  savedWorker,
			Image:   resourceType.Image,
			Version: resourceType.Version,
			BaseResourceType: &BaseResourceType{
				Name: resourceType.Type,
			},
		}

		uwrt, err := workerResourceType.FindOrCreate(tx, resourceType.UniqueVersionHistory)
		if err != nil {
			return nil, err
		}

		workerBaseResourceTypeIDs = append(workerBaseResourceTypeIDs, uwrt.ID)
	}

	_, err = psql.Delete("worker_base_resource_types").
		Where(sq.Eq{
			"worker_name": atcWorker.Name,
		}).
		Where(sq.NotEq{
			"id": workerBaseResourceTypeIDs,
		}).
		RunWith(tx).
		Exec()
	if err != nil {
		return nil, err
	}

	if atcWorker.CertsPath != nil {
		_, err := WorkerResourceCerts{
			WorkerName: atcWorker.Name,
			CertsPath:  *atcWorker.CertsPath,
		}.FindOrCreate(tx)
		if err != nil {
			return nil, err
		}
	}

	return savedWorker, nil
}
