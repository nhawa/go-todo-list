package presenter

import (
	"go-todo-list/lib"
	"net/http"
)

var (
	ErrLoginFailMessage = "Failed to log in Customer"
	ErrLoginFailCode    = 300201
)

type (
	AuthenticationPresenterInterface interface {
		SendLoginSuccessResponse(authenticationData AuthenticationData) lib.Response
		SendLoginFailResponse() lib.Response
	}

	AuthenticationData struct {
		AccessToken string `json:"access_token"`
		CustomerName string `json:"name"`
	}
	AuthenticationPresenter struct {
	}
)

func NewAuthenticationPresenter() AuthenticationPresenterInterface {
	return &AuthenticationPresenter{}
}

func (presenter *AuthenticationPresenter) SendLoginSuccessResponse(authenticationData AuthenticationData) lib.Response {
	response := lib.CreateResponse(200, 3020, "Success", authenticationData)
	return response
}

func (presenter *AuthenticationPresenter) SendLoginFailResponse() lib.Response {
	response := lib.CreateFailResponse(
		http.StatusUnauthorized,
		ErrLoginFailCode,
		ErrLoginFailMessage,
		nil,
		"Email atau kata sandi tidak tepat",
	)
	return response
}
