package routers

import (
	"github.com/FernandoCagale/c4-active/api/handlers"
	"github.com/gorilla/mux"
)

type SystemRoutes struct {
	healthHandler *handlers.HealthHandler
	activeHandler *handlers.ActiveHandler
}

func (routes *SystemRoutes) MakeHandlers() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", routes.healthHandler.Health).Methods("GET")
	r.HandleFunc("/v1/api/active/{id}", routes.activeHandler.FindByID).Methods("GET")
	r.HandleFunc("/v1/api/active/{id}", routes.activeHandler.UpdateByID).Methods("PUT")
	r.HandleFunc("/v1/api/active/{id}", routes.activeHandler.DeleteByID).Methods("DELETE")
	r.HandleFunc("/v1/api/active", routes.activeHandler.FindAll).Methods("GET")
	r.HandleFunc("/v1/api/active", routes.activeHandler.Create).Methods("POST")

	return r
}

func NewSystem(healthHandler *handlers.HealthHandler, activeHandler *handlers.ActiveHandler) *SystemRoutes {
	return &SystemRoutes{
		healthHandler: healthHandler,
		activeHandler: activeHandler,
	}
}
