package server

import (
	"social-backend/config"

	"github.com/labstack/echo/v4"
)

type Server struct {
	echo           *echo.Echo
}

func NewServer(echo *echo.Echo) *Server {
	return &Server{
		echo: echo,
	}
}

func(s *Server) Start() {
	e := s.echo
	e.Logger.Fatal(e.Start(":" + config.App().Port))
}
