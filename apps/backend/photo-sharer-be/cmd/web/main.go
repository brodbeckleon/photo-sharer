package main

import (
	"log"
	"net/http"
	"os"

	"example.com/m/v2/internal/api"
	"example.com/m/v2/internal/config"
)

func main() {
	// For demonstration. Use environment variables in production.
	os.Setenv("SECRET_CODE", "my-secret-code-123")
	os.Setenv("JWT_SECRET", "a-very-secret-key")

	cfg := config.Load()
	app := &api.Application{
		Config: cfg,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/login", app.LoginHandler)
	mux.HandleFunc("/api/protected", app.AuthMiddleware(app.ProtectedHandler))

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
