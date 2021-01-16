package router

import (
	"go-task/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/editor/graph/detailed/{id}", middleware.GetGraph).Methods("GET", "OPTIONS")
	router.HandleFunc("/editor/graph/detailed", middleware.GetAllGraphDetailed).Methods("GET", "OPTIONS")
	router.HandleFunc("/editor/graph/nondetailed", middleware.GetAllGraphNonDetailed).Methods("GET", "OPTIONS")
	router.HandleFunc("/editor/graph", middleware.CreateGraph).Methods("POST", "OPTIONS")
	router.HandleFunc("/editor/graph/{id}", middleware.UpdateGraph).Methods("PUT", "OPTIONS")
	router.HandleFunc("/editor/graph/{id}", middleware.DeleteGraph).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/editor/patrol/{id}", middleware.GetPatrol).Methods("GET", "OPTIONS")
	router.HandleFunc("/editor/patrol", middleware.GetAllPatrol).Methods("GET", "OPTIONS")
	router.HandleFunc("/editor/patrol", middleware.CreatePatrol).Methods("POST", "OPTIONS")
	router.HandleFunc("/editor/patrol/{id}", middleware.UpdatePatrol).Methods("PUT", "OPTIONS")
	router.HandleFunc("/editor/patrol/{id}", middleware.DeletePatrol).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/editor/robot", middleware.GetAllRobots).Methods("GET", "OPTIONS")
	router.HandleFunc("/editor/robot", middleware.CreateRobot).Methods("POST", "OPTIONS")
	router.HandleFunc("/editor/robot/{id}", middleware.DeleteRobot).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/editor/task", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/editor/task/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/editor/task/patrol", middleware.GetAllPatrolTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/editor/task/goto", middleware.GetAllGotoTasks).Methods("GET", "OPTIONS")

	router.HandleFunc("/editor/collection", middleware.CreateCollection).Methods("POST", "OPTIONS")
	router.HandleFunc("/editor/collection", middleware.GetAllCollection).Methods("GET", "OPTIONS")
	router.HandleFunc("/editor/collection/{id}", middleware.DeleteCollection).Methods("DELETE", "OPTIONS")

	return router
}
