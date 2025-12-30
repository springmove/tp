package base

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	DefaultPageSize = 20
)

type IQuery interface {
	FromCtx(ctx echo.Context) IQuery
	loadDB(db *gorm.DB)
	ToQuery() *gorm.DB
	ToURLQueryString() string
}

type QueryBase struct {
	IQuery

	db             *gorm.DB
	urlQueryString string

	Paging   bool
	Page     int64
	PageSize int64
	IDs      []string
}

func (s *QueryBase) ToURLQueryString() string {
	s.urlQueryString = s.urlQueryString + fmt.Sprintf("Page=%d&PageSize=%d&IDs=%s",
		s.Page,
		s.PageSize,
		strings.Join(s.IDs, ","))

	return s.urlQueryString
}

func (s *QueryBase) loadDB(db *gorm.DB) {
	s.db = db
}

func (s *QueryBase) FromCtx(ctx echo.Context) IQuery {
	ids := ctx.QueryParam("IDs")
	if ids != "" {
		s.IDs = strings.Split(ids, ",")
	}

	page, err := strconv.ParseInt(ctx.QueryParam("Page"), 10, 32)
	if err != nil {
		page = 0
	}

	s.Page = page

	pageSize, err := strconv.ParseInt(ctx.QueryParam("PageSize"), 10, 32)
	if err != nil {
		pageSize = DefaultPageSize
	}

	s.PageSize = pageSize

	return s
}

func (s *QueryBase) ToQuery() *gorm.DB {
	q := s.db.Where("deleted = ?", false)

	if len(s.IDs) > 0 {
		q = q.Where("id in (?)", s.IDs)
	}

	if s.Paging {
		if s.PageSize == 0 {
			s.PageSize = DefaultPageSize
		}

		q = q.Limit(int(s.PageSize)).Offset(int(s.Page * s.PageSize))
	}

	return q
}

func CreateQueryFromContext[T IQuery](query T, db *gorm.DB, ctx ...echo.Context) T {
	query.loadDB(db)

	if len(ctx) > 0 {
		query.FromCtx(ctx[0])
	}

	return query
}

func QueryModels[T IQuery, I ISerialize[I]](query T) ([]I, error) {

	models := []I{}

	listQuery := query.ToQuery()
	if err := listQuery.Find(&models).Error; err != nil {
		return nil, err
	}

	return SerializeModels(models), nil
}
