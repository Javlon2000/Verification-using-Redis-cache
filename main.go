package main

import (
	"os"
	"log"
	"net/http"

	c "github.com/Javlon2000/Verification-using-Redis-database/controllers"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Cannot load the .env: %v", err)
	}

	port := os.Getenv("APP_HTTP_PORT")

	http.HandleFunc("/signup", c.SignUP)
	http.HandleFunc("/verify", c.Verify)

	log.Println("Listening on the localhost", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}