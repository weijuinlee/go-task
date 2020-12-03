package main

import (
	"fmt"
	"go-task/router"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	r := router.Router()

	fmt.Println("Starting server on the port 8080...")

	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r))
	// log.Fatal(http.ListenAndServe(":8080", r))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)))
}
