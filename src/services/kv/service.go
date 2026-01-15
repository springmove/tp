package kv

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/springmove/sptty"
	"github.com/springmove/tp/src/base"
)

type Service struct {
	sptty.BaseService

	kv base.IClientContext

	adminToken string
}

func (s *Service) Init(app sptty.ISptty) error {

	if base.IGoRedis == nil {
		return fmt.Errorf(base.ErrIGoRedisIsRequired)
	}

	if base.IGorm == nil {
		return fmt.Errorf(base.ErrIGormIsRequired)
	}

	if base.IEcho == nil {
		return fmt.Errorf(base.ErrIEchoIsRequired)
	}

	s.kv = base.IGoRedis.ClientContext()

	if err := base.IGorm.AddModels([]any{
		&base.KV{},
	}); err != nil {
		return err
	}

	// 查询kv
	base.IEcho.Srv().GET("/kv", s.routeGetKV)

	// 更新kv缓存（后台操作）
	base.IEcho.Srv().POST("/kv-cache", s.routePostAdminKVCache)

	return nil
}

func (s *Service) ServiceName() string {
	return base.ServiceKV
}

func (s *Service) SetAdminToken(token string) {
	s.adminToken = token
}

func (s *Service) GetValueByKey(key string) (*base.KV, error) {

	kv := base.KV{}

	val, err := s.kv.Client().Get(*s.kv.Ctx(), key).Result()
	if err != nil {
		if strings.Contains(err.Error(), "redis: nil") {
			return kv.Serialize(), nil
		} else {
			return nil, err
		}
	}

	if err := json.Unmarshal([]byte(val), &kv); err != nil {
		return nil, err
	}

	return kv.Serialize(), nil
}

func (s *Service) UpdateKVCache(token string) error {

	if token != s.adminToken {
		return fmt.Errorf(base.ErrAccessDenied)
	}

	kvs, err := base.QueryModels[*base.KV](base.CreateQueryFromContext(&base.QueryKV{
		QueryBase: base.QueryBase{
			Paging: false,
		},
	}, base.IGorm.DB()))

	if err != nil {
		return err
	}

	if len(kvs) == 0 {
		return nil
	}

	tm := map[string]any{}
	for _, v := range kvs {
		val, err := json.Marshal(v)
		if err == nil {
			tm[v.ID] = string(val)
		}
	}

	if err := s.kv.Client().MSet(*s.kv.Ctx(), tm).Err(); err != nil {
		return err
	}

	sptty.Log(sptty.InfoLevel, "Update KV Cache Successfully", s.ServiceName())

	return nil
}
