package base

import "github.com/labstack/echo/v4"

const (
	ServiceEcho = "echo"
)

type IServiceEcho interface {
	Srv() *echo.Echo
}
