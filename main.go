package main

import (
	"fmt"
	"log" // Added to handle server errors
	"math/rand"
	"net/http"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Random Number API! Use /random to get a random number.")
}

func randomNumber(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	number := rand.Intn(100)         // Generate a random number (0-99)
	fmt.Fprintf(w, "Random Number: %d", number)
}

func main() {
	http.HandleFunc("/", home)               // Define root handler
	http.HandleFunc("/random", randomNumber) // Define /random handler
	fmt.Println("Server is running on port 8080")

	// Handle server error using log.Fatal
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
