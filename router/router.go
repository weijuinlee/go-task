package router

import (
	"go-task/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/editor/patrol/{id}", middleware.GetPatrol).Methods("GET", "OPTIONS")
	router.HandleFunc("/editor/patrol", middleware.GetAllPatrol).Methods("GET", "OPTIONS")
	router.HandleFunc("/editor/patrol", middleware.CreatePatrol).Methods("POST", "OPTIONS")
	router.HandleFunc("/editor/patrol/{id}", middleware.UpdatePatrol).Methods("PUT", "OPTIONS")
	router.HandleFunc("/editor/patrol/{id}", middleware.DeletePatrol).Methods("DELETE", "OPTIONS")

	return router
}
