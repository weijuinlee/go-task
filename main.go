package main

import (
	"fmt"
	"go-task/router"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	r := router.Router()

	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:8080"), handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)))
}
