package main

import (
	"log"
	"net/http"

	"github.com/student-management-service/controllers"
)

func main() {
	// Test comment
	log.Println("Wellcome To School Management Service")
	controllers.RegisterController()
	http.ListenAndServe(":5000", nil)
}
