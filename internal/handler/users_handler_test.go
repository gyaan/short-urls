package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gyaan/short-urls/internal/models"
	"github.com/gyaan/short-urls/mocks"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func Test_userHandler_RegisterUser(t *testing.T) {
	//request data
	requestData, err := json.Marshal(CreateUserRequest{
		Name:            "Gyaan",
		Email:           "test@gmail.com",
		Password:        "test1223",
		ConfirmPassword: "test1223",
	})

	//user details
	user := models.User{
		Name:     "Gyaan",
		Email:    "test@gmail.com",
		Password: "password hash",
		Status:   1,
	}

	if err != nil {
		t.Fatal(err)
	}

	//create request
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(requestData))
	if err != nil {
		t.Fatal(err)
	}

	//creat new response writer
	rr := httptest.NewRecorder()

	//user repository mock
	userRepository := mocks.Users{}
	userRepository.On("CreateUser", context.TODO(), mock.Anything).Return(&user, nil)
	newUserHandler := NewUserHandler(&userRepository)
	handler := http.HandlerFunc(newUserHandler.RegisterUser)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func Test_userHandler_UpdateUser(t *testing.T) {

	//request data
	requestData, err := json.Marshal(UpdateUserRequest{
		Email:    "test@gmail.com",
		Password: "test1223",
		Status:   0,
	})

	if err != nil {
		t.Fatal(err)
	}

	//create request
	req, err := http.NewRequest("PUT", "/users/5e50127d5d6c5e5e30dd6e79", bytes.NewBuffer(requestData))
	if err != nil {
		t.Fatal(err)
	}

	//user details
	user := models.User{
		Name:     "Gyaan",
		Email:    "test@gmail.com",
		Password: "password hash",
		Status:   1,
	}

	//creat new response writer
	rr := httptest.NewRecorder()

	//user repository mock
	userRepository := mocks.Users{}
	userRepository.On("GetUserDetailsById", context.TODO(), mock.Anything).Return(&user, nil)
	userRepository.On("UpdateUser", context.TODO(), mock.Anything, mock.Anything).Return(nil)
	newUserHandler := NewUserHandler(&userRepository)
	handler := http.HandlerFunc(newUserHandler.UpdateUser)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func Test_userHandler_GetUser(t *testing.T) {
	//create request
	req, err := http.NewRequest("GET", "/users", bytes.NewBufferString(url.Values{"user_id": {`xyx`}}.Encode()))
	if err != nil {
		t.Fatal(err)
	}

	//creat new response writer
	rr := httptest.NewRecorder()

	//user details
	user := models.User{
		Name:     "Gyaan",
		Email:    "test@gmail.com",
		Password: "password hash",
		Status:   1,
	}

	//user repository mock
	usersRepository := mocks.Users{}
	usersRepository.On("GetUserDetailsById", context.TODO(), mock.Anything).Return(&user, nil)

	newShortUrlHandler := NewUserHandler(&usersRepository)
	handler := http.HandlerFunc(newShortUrlHandler.GetUser)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
