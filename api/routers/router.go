package routers

import (
	"github.com/FernandoCagale/c4-active/config"
	"github.com/FernandoCagale/c4-active/pkg/active"
	"github.com/FernandoCagale/c4-active/pkg/infra/infra"
	"github.com/gorilla/mux"
)

//NewRouter infra
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.NotFoundHandler = infra.NotFoundHandler()
	router.MethodNotAllowedHandler = infra.NotAllowedHandler()
	return router
}

//makeGorm database postgres
func makeGorm(env *config.Config) active.UseCase {
	return active.NewService(active.NewGormRepository(env))
}
