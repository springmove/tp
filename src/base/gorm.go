package base

import (
	"time"

	"github.com/springmove/sptty"
	"gorm.io/gorm"
)

const (
	ServiceGorm = "gorm"

	DBPostgres = "postgres"
	DBMysql    = "mysql"

	DefaultMysqlCharset = "utf8mb4"
)

const (
	ErrTypeNotSupported = "ErrTypeNotSupported"
)

var IGorm IServiceGorm

type IServiceGorm interface {
	DB(index ...int) *gorm.DB
}

type BaseModel struct {
	ID        string `gorm:"size:32;primary_key"`
	Deleted   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *BaseModel) Init() *BaseModel {
	if s.ID == "" {
		s.ID = sptty.GenerateUID()
	}

	s.CreatedAt = time.Now().UTC()
	s.UpdatedAt = time.Now().UTC()
	s.Deleted = false

	return s
}

func (s *BaseModel) Serialize() *BaseModel {
	s.CreatedAt = s.CreatedAt.UTC()
	s.UpdatedAt = s.UpdatedAt.UTC()

	return s
}
