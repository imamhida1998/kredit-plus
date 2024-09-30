package repo

import (
	"context"
	"kredit-plus/service/model"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func NewRedisRepository(session *redis.Client) RedisRepository {
	return &redisRepositoryImpl{
		Session: *session,
	}
}

type redisRepositoryImpl struct {
	Session redis.Client
}

func (repository *redisRepositoryImpl) StoreValue(params *model.RedisStoreRequest) error {

	err := repository.Session.Set(ctx, params.KeyValue, []byte(params.Value), time.Duration(params.Lifetime)*time.Second).Err()
	if err != nil {
		return err
	}

	return nil

}

func (repository *redisRepositoryImpl) GetValue(keyValue string) *model.RedisValueEntity {

	storage, _ := repository.Session.Get(ctx, keyValue).Bytes()

	result := &model.RedisValueEntity{
		Value: string(storage),
	}
	//End Track Event Appd

	return result

}

func (repository *redisRepositoryImpl) DelValue(keyValue string) error {

	err := repository.Session.Del(ctx, keyValue).Err()
	if err != nil {
		return err
	}
	return nil

}

func (repository *redisRepositoryImpl) GetTtl(keyValue string) time.Duration {
	ttl, _ := repository.Session.TTL(ctx, keyValue).Result()
	return ttl
}
