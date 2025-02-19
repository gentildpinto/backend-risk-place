package server

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/risk-place-angola/backend-risk-place/api/rest/dependency"
	"github.com/risk-place-angola/backend-risk-place/infra/db/drive/postgres"
)

type Server struct {
	Router *echo.Echo
}

func NewServer(router *echo.Echo) *Server {
	return &Server{
		Router: router,
	}
}

func (server *Server) Start() {

	db, err := postgres.ConnectionPostgres()
	if err != nil {
		log.Panicln(err)
	}

	defer db.Close()

	dependency.Dependency(db, server.Router)

	server.Router.Logger.Fatal(server.Router.Start(":8000"))
}
