package application

import (
	_loginController "go-todo-list/authentication/controller"
	_authRepository "go-todo-list/authentication/repository"
	_authService "go-todo-list/authentication/service"
	_todoController "go-todo-list/todo/controller"
	_todoRepository "go-todo-list/todo/repository"
)

func (application *Application) register() {

	// middleware
	//basicMiddleware := _middleware.New()
	// repository
	todoRepository := _todoRepository.NewTodoRepository(application.ApplicationDB)
	userRepository := _authRepository.NewUserRepository(application.ApplicationDB)
	tokenRepository := _authRepository.NewTokenRepository(application.ApplicationDB)
	//service
	tokenService := _authService.NewTokenService(tokenRepository)
	authService := _authService.New(userRepository, tokenService)
	// Delivery layer / HTTP (Controller)
	_loginController.New(application.mux, authService)
	_todoController.NewTodoController(application.mux, todoRepository)
}
