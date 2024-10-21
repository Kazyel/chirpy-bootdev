package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Kazyel/chirpy-bootdev/internal/auth"
	"github.com/Kazyel/chirpy-bootdev/internal/database"
	"github.com/Kazyel/chirpy-bootdev/utils"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func (cfg *ApiConfig) HandlerCreateUsers(w http.ResponseWriter, r *http.Request) {
	type userRequest struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	body := &userRequest{}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		utils.RespondWithError(w, 403, "Something went wrong")
		return
	}

	hashedPassword, err := auth.HashPassword(body.Password)

	if err != nil {
		utils.RespondWithError(w, 403, "Something went wrong")
		return
	}

	newUser, err := cfg.db.CreateUser(r.Context(), database.CreateUserParams{
		Email:          body.Email,
		HashedPassword: hashedPassword,
	})

	if err != nil {
		utils.RespondWithError(w, 500, err.Error())
		return
	}

	userResponse := UserResponse{
		ID:        newUser.ID,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
		Email:     newUser.Email,
	}

	marshalResponse, err := json.Marshal(userResponse)
	if err != nil {
		utils.RespondWithError(w, 500, err.Error())
		return
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(201)
	w.Write(marshalResponse)
}
