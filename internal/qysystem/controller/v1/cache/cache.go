package cache

import (
	"context"
	"fmt"
	pb "gitee.com/windcoder/qingyucms/internal/pkg/qy-api/proto/qysystem/v1"
	log "gitee.com/windcoder/qingyucms/internal/pkg/qy-log"
	"gitee.com/windcoder/qingyucms/internal/qysystem/store"
	"sync"
)

type Cache struct {
	store store.Factory
}

var (
	cacheServer *Cache
	once        sync.Once
)

func GetCacheInsOr(store store.Factory) (*Cache, error) {
	if store != nil {
		once.Do(func() {
			cacheServer = &Cache{store}
		})
	}

	if cacheServer == nil {
		return nil, fmt.Errorf("got nil cache server")
	}

	return cacheServer, nil
}

func (c Cache) ListSecrets(ctx context.Context, r *pb.ListSecretsRequest) (*pb.ListSecretsResponse, error) {
	log.L(ctx).Info("list secrets function called.")
	//opts := metav1.ListOptions{
	//	Offset: r.Offset,
	//	Limit:  r.Limit,
	//}
	//
	//secrets, err :=c.store.S
	panic("implement me")
}

func (c Cache) ListPolicies(ctx context.Context, request *pb.ListPoliciesRequest) (*pb.ListPoliciesResponse, error) {
	//TODO implement me
	panic("implement me")
}
