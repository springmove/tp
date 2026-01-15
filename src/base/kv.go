package base

import "gorm.io/gorm"

const (
	ServiceKV = "kv"
)

const (
	ErrIGormIsRequired    = "ErrIGormIsRequired"
	ErrIGoRedisIsRequired = "ErrIGoRedisIsRequired"
	ErrIEchoIsRequired    = "ErrIEchoIsRequired"
	ErrAccessDenied       = "ErrAccessDenied"
)

var IKV IServiceKV

type IServiceKV interface {
	SetAdminToken(token string)
	GetValueByKey(key string) (*KV, error)
	UpdateKVCache(token string) error
}

type KV struct {
	BaseModel

	Value string
}

func (s *KV) Serialize() *KV {
	s.BaseModel.Serialize()

	return s
}

type QueryKV struct {
	QueryBase
}

func (s *QueryKV) ToQuery() *gorm.DB {
	q := s.QueryBase.ToQuery()

	q = q.Order("created_at desc")

	return q
}
