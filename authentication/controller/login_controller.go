package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-todo-list/lib"
	"net/http"
)

type (
	LoginService interface {
		Login(loginData LoginData) lib.Response
	}
	LoginData struct {
		Username string
		Password string
		FCMToken string
	}
	LoginController struct {
		loginService LoginService
	}
)

func New(
	mux *mux.Router,
	loginService LoginService,
	)  {
	loginController := &LoginController{}
	loginController.loginService = loginService

	mux.HandleFunc("/customer/login", loginController.login).Methods("POST")
}

func (lc *LoginController) login(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	reqBody := json.NewDecoder(r.Body)
	var loginData LoginData
	err := reqBody.Decode(&loginData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	lc.loginService.Login(loginData).JSON(w)
}
