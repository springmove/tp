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
	FromCtx(ctx echo.Context)
	loadDB(db *gorm.DB)
	ToQuery(paging bool) *gorm.DB
	ToURLQueryString() string
}

type QueryBase struct {
	IQuery

	db             *gorm.DB
	urlQueryString string

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

func (s *QueryBase) FromCtx(ctx echo.Context) {
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
}

func (s *QueryBase) ToQuery(paging bool) *gorm.DB {
	q := s.db

	q = q.Where("deleted = ?", false)

	if s.PageSize == 0 {
		s.PageSize = DefaultPageSize
	}

	if len(s.IDs) > 0 {
		q = q.Where("id in (?)", s.IDs)
	}

	if paging {
		q = q.Limit(int(s.PageSize)).Offset(int(s.Page * s.PageSize))
	}

	return q
}

func CreateQueryFromContext(query IQuery, db *gorm.DB, ctx ...echo.Context) IQuery {
	query.loadDB(db)

	if len(ctx) > 0 {
		query.FromCtx(ctx[0])
	}

	return query
}
