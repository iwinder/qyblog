package base

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/go-redis/redis/v8"
	log "github.com/iwinder/qyblog/internal/pkg/logger"
	genericoptions "github.com/iwinder/qyblog/internal/pkg/options/base"
	"strconv"
	"sync"
	"time"
)

type RedisStore struct {
	Client *redis.Client
}

func (s *RedisStore) Close() {
	err := s.Client.Close()
	if err != nil {
		log.Fatalf("Close Redis error: %v", err)
	}
}

var (
	redisStore *RedisStore
	redisOnce  sync.Once
)

func GetRedisClientOr(opts *genericoptions.RedisOptions) (*RedisStore, error) {
	timeout := 5 * time.Second
	if opts.Timeout > 0 {
		timeout = time.Duration(opts.Timeout) * time.Second
	}
	poolSize := 10
	if opts.MaxActive > 0 {
		poolSize = opts.MaxActive
	}
	var tlsConfig *tls.Config

	if opts.UseSSL {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: opts.SSLInsecureSkipVerify,
		}
	}
	var err error
	var redisCliDB *redis.Client
	redisOnce.Do(func() {
		redisCliDB = redis.NewClient(&redis.Options{
			Addr:         opts.Host + ":" + strconv.Itoa(opts.Port),
			Password:     opts.Password,
			DB:           opts.Database,
			ReadTimeout:  timeout,
			WriteTimeout: timeout,
			DialTimeout:  timeout,
			IdleTimeout:  240 * timeout,
			PoolSize:     poolSize,
			TLSConfig:    tlsConfig,
		})
		redisStore = &RedisStore{Client: redisCliDB}
	})
	if redisStore.Client == nil {
		return nil, fmt.Errorf("failed to get redis store fatory, redisStore: %+v, error: %w", redisStore, err)
	}
	pingTimeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()
	err = redisCliDB.Ping(pingTimeout).Err()
	if err != nil {
		return nil, fmt.Errorf("redis connect error: %v", err)
	}
	return redisStore, nil
}

func getRedisAddrs(opts *genericoptions.RedisOptions) (addrs []string) {
	if len(opts.Addrs) != 0 {
		addrs = opts.Addrs
	}

	if len(addrs) == 0 && opts.Port != 0 {
		addr := opts.Host + ":" + strconv.Itoa(opts.Port)
		addrs = append(addrs, addr)
	}

	return addrs
}
