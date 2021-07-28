package app

import (
	"log"
	"mvc/controllers"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
