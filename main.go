package main

import (
	"fmt"
	"go-with-postgres/router"
	"log"
	"net/http"
)

func main() {
	rtr := router.Router()

	fmt.Println("Starting server on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", rtr))
}
