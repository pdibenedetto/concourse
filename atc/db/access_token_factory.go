package db

import (
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . AccessTokenFactory
type AccessTokenFactory interface {
	CreateAccessToken(token string, claims Claims) error
	GetAccessToken(token string) (AccessToken, bool, error)
	DeleteAccessToken(token string) error
}

func NewAccessTokenFactory(conn DbConn) AccessTokenFactory {
	return &accessTokenFactory{conn}
}

type accessTokenFactory struct {
	conn DbConn
}

func (a *accessTokenFactory) CreateAccessToken(token string, claims Claims) error {
	var expiry int64
	if claims.Expiry != nil {
		expiry = int64(*claims.Expiry)
	}
	_, err := psql.Insert("access_tokens").
		Columns("token", "sub", "expires_at", "claims").
		Values(token, claims.Subject, time.Unix(expiry, 0), claims).
		RunWith(a.conn).
		Exec()
	if err != nil {
		return err
	}
	return nil
}

func (a *accessTokenFactory) DeleteAccessToken(token string) error {
	_, err := psql.Delete("access_tokens").
		Where(sq.Eq{"token": token}).
		RunWith(a.conn).
		Exec()
	if err != nil {
		return err
	}
	return nil
}

func (a *accessTokenFactory) GetAccessToken(token string) (AccessToken, bool, error) {
	row := psql.Select("token", "claims").
		From("access_tokens").
		Where(sq.Eq{"token": token}).
		RunWith(a.conn).
		QueryRow()

	var accessToken AccessToken
	err := scanAccessToken(&accessToken, row)
	if err != nil {
		if err == sql.ErrNoRows {
			return AccessToken{}, false, nil
		}
		return AccessToken{}, false, err
	}
	return accessToken, true, nil
}
