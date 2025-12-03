package base

import "gorm.io/gorm"

const (
	ServiceGorm = "gorm"

	DBPostgres = "postgres"
	DBMysql    = "mysql"

	DefaultMysqlCharset = "utf8mb4"
)

const (
	ErrTypeNotSupported = "ErrTypeNotSupported"
)

type IServiceGorm interface {
	DB(index ...int) *gorm.DB
}
