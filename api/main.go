package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	router "github.com/omjogani/bookmark-manager/routes"
)

func main() {
	fmt.Println("Welcome to Bookmark Manager")

	// establishing router
	r := router.Router()
	fmt.Println("Server is getting started...")

	// Get PORT from .env
	envError := godotenv.Load(".env")
	if envError != nil {
		fmt.Println("ERROR: ", envError)
		log.Fatal("Could not load Environment Variables...")
	}
	port := os.Getenv("PORT_NO")
	fmt.Println(port)
	http.ListenAndServe(":4000", r)
	fmt.Println("Server is listening at port:", port)
}
