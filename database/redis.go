// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package database

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/libevm/common"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var _ RelayerDatabase = &RedisDatabase{}

type RedisDatabase struct {
	logger logging.Logger
	client *redis.Client
}

func NewRedisDatabase(logger logging.Logger, redisURL string, relayerIDs []RelayerID) (*RedisDatabase, error) {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		err = redactURLError(err)
		logger.Error(
			"Failed to parse Redis URL",
			zap.String("url", redactRedisURL(redisURL)),
			zap.Error(err),
		)
		return nil, err
	}

	// Create a new Redis client.
	// The server address, password, db index, and protocol version are extracted from the URL
	// If not provided in the URL, request timeouts use the default value of 3 seconds
	client := redis.NewClient(opts)
	return &RedisDatabase{
		logger: logger,
		client: client,
	}, nil
}

// redactRedisURL returns a form of the Redis URL that is safe to log. A Redis
// URL carries the connection password in its userinfo component, which
// config.Config tags sensitive:"true" so the startup config logger prints
// [REDACTED] for it; these helpers keep that guarantee on the parse error path.
// A URL that cannot be parsed is withheld entirely, since there is then no way
// to locate the credentials within it.
func redactRedisURL(redisURL string) string {
	u, err := url.Parse(redisURL)
	if err != nil {
		return "[REDACTED]"
	}
	return u.Redacted()
}

// redactURLError strips the raw URL out of a *url.Error, which redis.ParseURL
// returns unmodified when url.Parse fails and whose message embeds that URL
// verbatim, password included. The errors redis.ParseURL reports itself name
// only the offending component (scheme, database number, path, option), never
// the userinfo, so they are returned as-is.
func redactURLError(err error) error {
	var urlErr *url.Error
	if errors.As(err, &urlErr) {
		return fmt.Errorf("%s redis URL: %w", urlErr.Op, urlErr.Err)
	}
	return err
}

func (r *RedisDatabase) Get(relayerID common.Hash, key DataKey) ([]byte, error) {
	ctx := context.Background()
	compositeKey := constructCompositeKey(relayerID, key)
	val, err := r.client.Get(ctx, compositeKey).Result()
	if err != nil {
		r.logger.Debug("Error retrieving key from Redis",
			zap.String("key", compositeKey),
			zap.Error(err),
		)
		if err == redis.Nil {
			return nil, ErrKeyNotFound
		}
		return nil, err
	}
	return []byte(val), nil
}

func (r *RedisDatabase) Put(relayerID common.Hash, key DataKey, value []byte) error {
	ctx := context.Background()
	compositeKey := constructCompositeKey(relayerID, key)

	// Persistently store the value in Redis
	err := r.client.Set(ctx, compositeKey, value, 0).Err()
	if err != nil {
		r.logger.Error("Error storing key in Redis",
			zap.String("key", compositeKey),
			zap.Error(err),
		)
		return err
	}
	return nil
}

func (r *RedisDatabase) Close() error {
	return r.client.Close()
}

func constructCompositeKey(relayerID common.Hash, key DataKey) string {
	const keyDelimiter = "-"
	return strings.Join([]string{relayerID.Hex(), key.String()}, keyDelimiter)
}
