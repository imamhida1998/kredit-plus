package repo

import (
	"kredit-plus/service/model"
	"time"
)

type RedisRepository interface {
	StoreValue(params *model.RedisStoreRequest) error
	GetValue(keyValue string) *model.RedisValueEntity
	DelValue(keyValue string) error
	GetTtl(keyValue string) time.Duration
}
