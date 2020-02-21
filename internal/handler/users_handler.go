package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gyaan/short-urls/internal/models"
	"github.com/gyaan/short-urls/internal/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
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

type UserHandler interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userRepository repositories.Users
}

func NewUserHandler(users repositories.Users) UserHandler {
	return &userHandler{
		userRepository: users,
	}
}

//RegisterUser creates a new user
func (h *userHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var createUserRequest CreateUserRequest
	errResponse := models.ErrorResponse{ErrorMessage: "error in registering user", Retry: false}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&createUserRequest)

	if err != nil {
		log.Printf("error with create user request %v", err)
		http.Error(w, errResponse.Error(), http.StatusBadRequest)
		return
	}

	if len(createUserRequest.Name) == 0 || len(createUserRequest.Password) == 0 || len(createUserRequest.Email) == 0 || len(createUserRequest.ConfirmPassword) == 0 {
		errResponse.ErrorMessage = "name, email, password and confirm password are required fields."
		http.Error(w, errResponse.Error(), http.StatusBadRequest)
		return
	}

	if createUserRequest.Password != createUserRequest.ConfirmPassword {
		errResponse.ErrorMessage = "password and confirm password should be same."
		http.Error(w, errResponse.Error(), http.StatusBadRequest)
		return
	}

	//get password hash
	password, err := bcrypt.GenerateFromPassword([]byte(createUserRequest.Password), bcrypt.MinCost)
	if err != nil {
		log.Printf("error getting  password hash %v", err)
		http.Error(w, errResponse.Error(), http.StatusInternalServerError)
		return
	}

	user, err := h.userRepository.CreateUser(r.Context(), models.User{
		ID:       primitive.NewObjectIDFromTimestamp(time.Now()),
		Name:     createUserRequest.Name,
		Email:    createUserRequest.Email,
		Password: string(password),
		Status:   1,
	})
	if err != nil {
		log.Printf("Error in create new user in database %v", err)
		http.Error(w, errResponse.Error(), http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(user)
	if err != nil {
		log.Printf("Error marshaling user details %v", err)
		http.Error(w, errResponse.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(bytes))
}

//UpdateUser updates existing user
func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUserRequest UpdateUserRequest
	errResponse := models.ErrorResponse{ErrorMessage: "error in updating user details", Retry: false}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updateUserRequest)

	if err != nil {
		log.Printf("Error with update user request %v", err)
		http.Error(w, errResponse.Error(), http.StatusBadRequest)
		return
	}

	userId := fmt.Sprintf("%v", r.Context().Value("user_id"))
	if len(userId) == 0 {
		log.Printf("user id isn't in context!")
		http.Error(w, errResponse.Error(), http.StatusBadRequest)
		return
	}

	//get password hash
	password, err := bcrypt.GenerateFromPassword([]byte(updateUserRequest.Password), bcrypt.MinCost)
	if err != nil {
		log.Printf("Error generating password hash %v", err)
		http.Error(w, errResponse.Error(), http.StatusInternalServerError)
		return
	}

	//lets say only password,status and email can be changes
	//get existing user details
	existingUser, err := h.userRepository.GetUserDetailsById(r.Context(), userId)

	if err != nil {
		log.Printf("Error getting exising user details %v", err)
		http.Error(w, errResponse.Error(), http.StatusInternalServerError)
		return
	}

	//update only there is email in request
	if len(updateUserRequest.Email) > 0 {
		existingUser.Email = updateUserRequest.Email
	}

	//update only if status is different
	if updateUserRequest.Status != existingUser.Status {
		existingUser.Status = updateUserRequest.Status
	}

	//update only if there is password
	if len(updateUserRequest.Password) > 0 {
		existingUser.Password = string(password)
	}

	err = h.userRepository.UpdateUser(r.Context(), userId, existingUser)
	if err != nil {
		log.Printf("Error updating user details for %s, %v", userId, err)
		http.Error(w, errResponse.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "successfully update user details")
}

//GetUser returns user details
func (h *userHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	userId := fmt.Sprintf("%v", r.Context().Value("user_id"))
	errResponse := models.ErrorResponse{ErrorMessage: "error in fetching user details", Retry: false}

	//get user details
	user, err := h.userRepository.GetUserDetailsById(r.Context(), userId)
	if err != nil {
		log.Printf("error fetching user details %v", err)
		http.Error(w, errResponse.Error(), http.StatusBadRequest)
		return
	}
	bytes, err := json.Marshal(user)
	if err != nil {
		log.Printf("Error marshaling user details, %v", err)
		http.Error(w, errResponse.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(bytes))
}
