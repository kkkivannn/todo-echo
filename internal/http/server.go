package http

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Server struct {
	host string
	port int
	e    *echo.Echo
}

func New(e *echo.Echo, host string, port int) *Server {
	return &Server{e: e, host: host, port: port}
}

func (s *Server) MustRun() {
	host := fmt.Sprintf("%s:%d", s.host, s.port)
	if err := s.e.Start(host); err != nil {
		panic(err)
	}
}

func (s *Server) Stop(ctx context.Context) {
	if err := s.e.Shutdown(ctx); err != nil {
		log.Error(" Can't stop server")
	}
}
