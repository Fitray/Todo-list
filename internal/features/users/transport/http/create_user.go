package users_transport_http

import (
	"encoding/json"
	"net/http"

	core_logger "github.com/Fitray/Todo-list/internal/core/logger"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

type CreateUserResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func (s *UsersHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	log.Debug("Got create user request")

	var request CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {

	}

	w.WriteHeader(http.StatusOK)
}
