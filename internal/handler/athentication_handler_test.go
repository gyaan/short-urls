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
	"testing"

	"github.com/gyaan/short-urls/internal/config"
)

func Test_authenticationHandler_GetAccessToken(t *testing.T) {

	//request data
	requestData, err := json.Marshal(AccessTokenRequest{
		Name:     "gyaan33",
		Password: "121212",
	})

	//user details
	user := models.User{
		Name:     "gyaan33",
		Password: "$2a$04$Af4EM2IrnsSaj4F7aPUSdedeesImJ6oL4ASvvy5/qK5vSD00reEWO",
	}

	if err != nil {
		t.Fatal(err)
	}

	//create request
	req, err := http.NewRequest("POST", "/access-token", bytes.NewBuffer(requestData))
	if err != nil {
		t.Fatal(err)
	}

	//creat new response writer
	rr := httptest.NewRecorder()
	conf := config.New()

	//user repository mock
	userRepository := mocks.Users{}
	userRepository.On("GetUserDetailsByName", context.TODO(), mock.Anything).Return(&user, nil)
	newAuthenticationHandler := NewAuthenticationHandler(&userRepository, conf)
	handler := http.HandlerFunc(newAuthenticationHandler.GetAccessToken)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
