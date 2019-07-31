package routers

import (
	"github.com/FernandoCagale/c4-active/api/handlers/active"
	"github.com/FernandoCagale/c4-active/config"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

//MakeHandlers make url handlers
func MakeHandlers(r *mux.Router, n negroni.Negroni, env *config.Config) {
	service := makeGorm(env)

	r.Handle("/api/v1/actives", n.With(
		negroni.Wrap(active.FindAll(service)),
	)).Methods("GET")

	r.Handle("/api/health", n.With(
		negroni.Wrap(active.Health(service)),
	)).Methods("GET")
}
