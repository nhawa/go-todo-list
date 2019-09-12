package application

import (
	_middleware "go-todo-list/authentication/middleware"
	_todoController "go-todo-list/todo/controller"
	_todoRepository "go-todo-list/todo/repository"
)

func (application *Application) register() {

	// middleware
	basicMiddleware := _middleware.New()
	// repository
	todoRepository := _todoRepository.NewTodoRepository(application.ApplicationDB)
	// Delivery layer / HTTP (Controller)
	_todoController.NewTodoController(application.mux, todoRepository, basicMiddleware)
}
