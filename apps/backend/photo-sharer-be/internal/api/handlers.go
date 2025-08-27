package api

import (
	"encoding/json"
	"net/http"

	"example.com/m/v2/internal/auth"
	"example.com/m/v2/internal/config"
)

type LoginPayload struct {
	Code string `json:"code"`
}

type Application struct {
	Config *config.AppConfig
}

func (app *Application) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var p LoginPayload
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if p.Code != app.Config.SecretCode {
		http.Error(w, "Invalid code", http.StatusUnauthorized)
		return
	}

	tokenString, expirationTime, err := auth.CreateToken(app.Config.JWTSecret)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}

func (app *Application) ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is protected content."))
}
