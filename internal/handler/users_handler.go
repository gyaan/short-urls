package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gyaan/short-urls/internal/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

//CreateUserRequest
type CreateUserRequest struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

//UpdateUserRequest
type UpdateUserRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	Status   int    `json:"status"`
}

//CreateUser creates a new user
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUserRequest CreateUserRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&createUserRequest)

	if err != nil {
		log.Println("Error with create user request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//todo validate all the required fields
	if createUserRequest.Password != createUserRequest.ConfirmPassword {
		err = errors.New("password and confirm password aren't same")
		log.Printf("Password and confirm password aren't same")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//get password hash
	password, err := bcrypt.GenerateFromPassword([]byte(createUserRequest.Password), bcrypt.MinCost)
	if err != nil {
		log.Printf("Error getting  password hash")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := h.userRepository.CreateUser(r.Context(), models.User{
		Name:     createUserRequest.Name,
		Email:    createUserRequest.Email,
		Password: string(password),
		Status:   1,
	})

	if err != nil {
		log.Printf("Error in create new user")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(user)
	if err != nil {
		log.Printf("Error marshaling user details")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(bytes))
}

//UpdateUser updates existing user
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUserRequest UpdateUserRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updateUserRequest)

	if err != nil {
		log.Printf("Error with update user request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r) //todo take this from jwt token

	//get password hash
	password, err := bcrypt.GenerateFromPassword([]byte(updateUserRequest.Password), bcrypt.MinCost)
	if err != nil {
		log.Printf("Error generating password hash")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//todo update only if field available
	err = h.userRepository.UpdateUser(r.Context(), vars["user_id"], models.User{
		Status:   updateUserRequest.Status,
		Password: string(password),
		Email:    updateUserRequest.Email,
	})

	if err != nil {
		log.Printf("Error updating user details for %s", vars["user_id"])
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "successfully update user details")
}
