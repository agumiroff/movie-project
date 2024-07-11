package httpserver

import (
	"context"
	"encoding/json"
	"movie-project/internal/app/domain"
	"net/http"
)

func (s *AppServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	user, err := domain.NewUser(domain.NewUserData{
		Username: "sss",
		Password: "sss",
		Admin:    false,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = s.userService.CreateUser(
		context.Background(),
		user,
	)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
