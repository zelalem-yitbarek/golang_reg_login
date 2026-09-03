package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Response{
		Message: "Hello from Go!",
		Status:  "OK",
	}

	json.NewEncoder(w).Encode(response)
}

func main() {

	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env:", err)
	}

	// Get MongoDB connection string
	mongoURI := os.Getenv("MONGO_URI")

	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set")
	}

	// Connect to MongoDB
	client, err := mongo.Connect(
		options.Client().ApplyURI(mongoURI),
	)
	if err != nil {
		log.Fatal("MongoDB client error:", err)
	}

	// Test MongoDB connection
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌ MongoDB connection failed:", err)
	}

	fmt.Println("✅ MongoDB connected successfully!")
	fmt.Println("Database: goLang")

	// Start API
	http.HandleFunc("/api/hello", helloHandler)

	fmt.Println("🚀 Server running on http://localhost:8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}