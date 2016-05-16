package main

import (
	"log"
	"net/http"

	"{{.GoPathPrefix}}/{{.AppName}}/middleware"
)

func main() {
	http.Handle("/", middleware.Handle(nil, middleware.PrintLog, middleware.ServeFilesFromDir("assets")))

	log.Println("Server is about to listen at port 8000.")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Printf("Could not start server at port 8000: %v\n", err)
	}
}
