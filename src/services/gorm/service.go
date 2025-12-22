package gorm

import (
	"fmt"

	"github.com/springmove/sptty"
	"github.com/springmove/tp/src/base"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service struct {
	sptty.BaseService

	cfg Config

	dbs []*gorm.DB
}

func (s *Service) ServiceName() string {
	return base.ServiceGorm
}

func (s *Service) Init(app sptty.ISptty) error {
	if err := app.GetConfig(s.ServiceName(), &s.cfg); err != nil {
		return err
	}

	if !s.cfg.Enable {
		sptty.Log(sptty.InfoLevel, "Service Disabled", s.ServiceName())
		return nil
	}

	if err := s.initDBs(); err != nil {
		return err
	}

	return nil
}

func (s *Service) DB(index ...int) *gorm.DB {
	targetDB := 0
	if len(index) > 0 {
		targetDB = index[0]
	}

	return s.dbs[targetDB]
}

func (s *Service) AddModels(models []any, db ...*gorm.DB) error {

	targetDB := s.DB()
	if len(db) > 0 {
		targetDB = db[0]
	}

	if err := targetDB.AutoMigrate(models); err != nil {
		return err
	}

	return nil
}

func (s *Service) initDBs() error {
	var db *gorm.DB
	var err error
	for _, v := range s.cfg.Configs {
		switch v.Type {
		case base.DBPostgres:
			db, err = gorm.Open(postgres.Open(getConnStr(&v)), &gorm.Config{})
			if err != nil {
				return err
			}

		case base.DBMysql:
			db, err = gorm.Open(mysql.Open(getConnStr(&v)), &gorm.Config{})
			if err != nil {
				return err
			}

		default:
			return fmt.Errorf(base.ErrTypeNotSupported)
		}

		s.dbs = append(s.dbs, db)
	}

	return nil
}

func getConnStr(cfg *DBConfig) string {
	connStr := ""

	switch cfg.Type {
	case base.DBPostgres:
		// ex: host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai
		connStr = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s connect_timeout=%d sslmode=disable",
			cfg.Host,
			cfg.Port,
			cfg.User,
			cfg.Name,
			cfg.Pwd,
			cfg.Timeout)

	case base.DBMysql:
		if cfg.Charset == "" {
			cfg.Charset = base.DefaultMysqlCharset
		}

		// ex: user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
		connStr = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			cfg.User,
			cfg.Pwd,
			cfg.Host,
			cfg.Port,
			cfg.Name,
			cfg.Charset)

	default:
		return connStr
	}

	return connStr
}
