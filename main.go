package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kasumusof/lannister-pay/pkg"
)

func init() {
	log.Println("Starting server...")
	log.Println("Listen on port:", pkg.Port)
	log.Println("Press CTRL+C to stop")
}

func main() {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", pkg.Port),
		Handler: pkg.Router(),
	}

	log.Fatal(server.ListenAndServe())
}
