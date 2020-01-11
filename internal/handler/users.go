package handler

import (
	"context"
	"fmt"
	"github.com/gyaan/short-urls/internal/models"
	"net/http"
)

//CreateUser creates a new user
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	user, err := h.userRepository.CreateUser(context.Background(), models.User{})
	fmt.Println(user)
	fmt.Println(err)
}

//UpdateUser updates existing user
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	user, err := h.userRepository.UpdateUser(context.Background(), models.User{})
	fmt.Println(user)
	fmt.Println(err)
}
