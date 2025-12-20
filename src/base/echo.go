package base

import "github.com/labstack/echo/v4"

const (
	ServiceEcho = "echo"
)

var IEcho IServiceEcho

type IServiceEcho interface {
	Srv() *echo.Echo
}
