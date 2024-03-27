package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	server := http.Server{
		Handler: router,
		Addr:    ":" + os.Getenv("SERVER_PORT"),
	}

	fmt.Println("Server is running on port 8080")
	server.ListenAndServe()
}
