package echo

import (
	"fmt"

	v4 "github.com/labstack/echo/v4"
	"github.com/springmove/sptty"
	"github.com/springmove/tp/src/base"
)

type Service struct {
	sptty.BaseService

	cfg Config
	srv *v4.Echo
}

func (s *Service) ServiceName() string {
	return base.ServiceEcho
}

func (s *Service) Init(app sptty.ISptty) error {

	if err := sptty.GetApp().GetConfig(s.ServiceName(), &s.cfg); err != nil {
		return err
	}

	if !s.cfg.Enable {
		sptty.Log(sptty.InfoLevel, "Service Disabled", s.ServiceName())
		return nil
	}

	go func() {
		if err := s.Srv().Start(s.cfg.Port); err != nil {
			sptty.Log(sptty.ErrorLevel, fmt.Sprintf("Echo Server Err: %s", err.Error()), s.ServiceName())
			return
		}
	}()

	s.showRoutes()

	return nil
}

func (s *Service) Srv() *v4.Echo {
	if s.srv == nil {
		s.srv = v4.New()
	}

	return s.srv
}

func (s *Service) showRoutes() {
	routes := s.Srv().Routes()
	for _, v := range routes {
		fmt.Printf("%s %s\n", v.Method, v.Path)
	}
}
