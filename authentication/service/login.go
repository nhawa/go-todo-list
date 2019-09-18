package service

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go-todo-list/authentication/controller"
	"go-todo-list/authentication/presenter"
	"go-todo-list/authentication/repository"
	"go-todo-list/lib"
)

type (
	AuthenticationData = presenter.AuthenticationData
	LoginService struct {
		userRepository *repository.UserRepository
		tokenService TokenService
	}
)

func New(
	userRepository *repository.UserRepository,
	tokenService TokenService,
	) controller.LoginService{
	return &LoginService{
		userRepository,
		tokenService,
	}
	//loginService := &LoginService{}
	//loginService.userRepository = userRepository
	//
	//return loginService
}

func (loginService *LoginService) Login(loginData controller.LoginData) lib.Response {
	authPresenter := presenter.NewAuthenticationPresenter()

	user, err := loginService.userRepository.FindOneForCustomerLogin(loginData.Username)
	if err != nil {
		logrus.Error(err)
		return authPresenter.SendLoginFailResponse()
	}
	logrus.Info(user)
	isPasswordVerified := Verify(loginData.Password, user.Password)

	if !isPasswordVerified {
		fmt.Println("password not valid")
		return authPresenter.SendLoginFailResponse()
	}

	jwtToken, err := loginService.tokenService.CreateJWT(user.ID)

	if err != nil {
		fmt.Println(err.Error())
		return authPresenter.SendLoginFailResponse()
	}

	authenticationData := AuthenticationData{
		AccessToken:  jwtToken.AccessToken,
		//RefreshToken: jwtToken.RefreshToken,
		CustomerName: user.Name,
	}
	fmt.Println(authenticationData)
	return authPresenter.SendLoginSuccessResponse(authenticationData)
	//json.Marshal(authenticationData)
}