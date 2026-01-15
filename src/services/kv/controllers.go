package kv

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/springmove/sptty"
	"github.com/springmove/tp/src/base"
)

func (s *Service) routeGetKV(ctx echo.Context) error {
	var err error
	code := http.StatusBadRequest
	var body any

	body, err = s.GetValueByKey(ctx.QueryParam("key"))
	if err != nil {
		goto END
	}

	code = http.StatusOK

END:
	if err != nil {
		body = sptty.RequestError{
			Code: err.Error(),
			Msg:  err.Error(),
		}
	}

	return ctx.JSON(code, body)
}

func (s *Service) routePostAdminKVCache(ctx echo.Context) error {
	var err error
	code := http.StatusBadRequest
	var body any

	if err = s.UpdateKVCache(ctx.Request().Header.Get(base.HeaderAuth)); err != nil {
		goto END
	}

	code = http.StatusOK

END:
	if err != nil {
		body = sptty.RequestError{
			Code: err.Error(),
			Msg:  err.Error(),
		}
	}

	return ctx.JSON(code, body)
}
