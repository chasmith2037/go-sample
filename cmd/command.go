package cmd

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go-sample/config"
	"go-sample/routes"
	"net/http"
)

type Command struct {
	Router *chi.Mux
	Config *config.Config
}

func Initialize(config *config.Config) Command {
	command := Command{}

	command.Router = chi.NewRouter()
	command.Router.Use(middleware.Logger)

	command.Config = config
	command.setRoutes()

	return command
}

func (c *Command) setRoutes() {
	c.Router.Route("/sample", func(r chi.Router) {
		r.Mount("/api/v1", routes.GetRoutes())
	})
}

func (c *Command) Run() {
	port := ":" + c.Config.HostPort
	http.ListenAndServe(port, c.Router)
}
