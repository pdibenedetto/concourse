package creds

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type SecretCacheConfig struct {
	Enabled          bool          `long:"secret-cache-enabled" description:"Enable in-memory cache for secrets"`
	Duration         time.Duration `long:"secret-cache-duration" default:"1m" description:"If the cache is enabled, secret values will be cached for not longer than this duration (it can be less, if underlying secret lease time is smaller)"`
	DurationNotFound time.Duration `long:"secret-cache-duration-notfound" default:"10s" description:"If the cache is enabled, secret not found responses will be cached for this duration"`
	PurgeInterval    time.Duration `long:"secret-cache-purge-interval" default:"10m" description:"If the cache is enabled, expired items will be removed on this interval"`
}

type CachedSecrets struct {
	secrets     Secrets
	cacheConfig SecretCacheConfig
	cache       *cache.Cache
}

type CacheEntry struct {
	value      any
	expiration *time.Time
	found      bool
}

func NewCachedSecrets(secrets Secrets, cacheConfig SecretCacheConfig) *CachedSecrets {
	// Create a cache with:
	// - default expiration time for entries set to 'cacheConfig.Duration'
	// - purges expired items regularly, on every `cacheConfig.PurgeInterval` after creation
	return &CachedSecrets{
		secrets:     secrets,
		cacheConfig: cacheConfig,
		cache:       cache.New(cacheConfig.Duration, cacheConfig.PurgeInterval),
	}
}

func (cs *CachedSecrets) Get(secretPath string) (any, *time.Time, bool, error) {
	return cs.GetWithParams(secretPath, SecretLookupParams{})
}

func (cs *CachedSecrets) NewSecretLookupPaths(teamName string, pipelineName string, allowRootPath bool) []SecretLookupPath {
	return cs.NewSecretLookupPathsWithParams(SecretLookupParams{Team: teamName, Pipeline: pipelineName}, allowRootPath)
}

func (cs *CachedSecrets) GetWithParams(secretPath string, context SecretLookupParams) (any, *time.Time, bool, error) {
	// if there is a corresponding entry in the cache, return it
	entry, found := cs.cache.Get(secretPath)
	if found {
		result := entry.(CacheEntry)
		return result.value, result.expiration, result.found, nil
	}

	// otherwise, let's make a request to the underlying secret manager
	value, expiration, found, err := GetWithParams(cs.secrets, secretPath, context)

	// we don't want to cache errors, let the errors be retried the next time around
	if err != nil {
		return nil, nil, false, err
	}

	// here we want to cache secret value, expiration, and found flag too
	// meaning that "secret not found" responses will be cached too!
	entry = CacheEntry{value: value, expiration: expiration, found: found}

	if found {
		// take default cache ttl
		duration := cs.cacheConfig.Duration
		if expiration != nil {
			// if secret lease time expires sooner, make duration smaller than default duration
			// also if the duration is less than or equal to 0, use default duration (it would cache forever)
			itemDuration := time.Until(*expiration)
			if itemDuration < duration && itemDuration > 0 {
				duration = itemDuration
			}
		}
		cs.cache.Set(secretPath, entry, duration)
	} else {
		cs.cache.Set(secretPath, entry, cs.cacheConfig.DurationNotFound)
	}

	return value, expiration, found, nil
}

func (cs *CachedSecrets) NewSecretLookupPathsWithParams(context SecretLookupParams, allowRootPath bool) []SecretLookupPath {
	return NewSecretLookupPathsWithParams(cs.secrets, context, allowRootPath)
}
