package base

import "github.com/labstack/echo/v4"

const (
	ServiceEcho = "echo"
)

var Echo IServiceEcho

type IServiceEcho interface {
	Srv() *echo.Echo
}
