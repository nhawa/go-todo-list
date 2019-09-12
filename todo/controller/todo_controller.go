package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-todo-list/authentication/middleware"
	"go-todo-list/todo/model"
	"go-todo-list/todo/presenter"
	"go-todo-list/todo/repository"
	"net/http"
	"strconv"
)

type (
	TodoController struct {
		todoRepo         repository.TodoRepository
	}
)


func NewTodoController(
	mux *mux.Router,
	todoRepo repository.TodoRepository,
	basicMiddleware middleware.BasicMiddlewareInterface,
) *TodoController {

	todoController := &TodoController{}
	todoController.todoRepo = todoRepo

	getTodoListHandler := http.HandlerFunc(todoController.getTodoList)
	mux.HandleFunc("/", todoController.getTodoList).Methods("GET")
	mux.Handle("/todo-lists", basicMiddleware.Middleware(getTodoListHandler)).Methods("GET")
	mux.HandleFunc("/todo-lists", todoController.postTodoList).Methods("POST")
	mux.HandleFunc("/todo-lists/{id}", todoController.getTodoListDetail).Methods("GET")
	mux.HandleFunc("/todo-lists/{id}", todoController.updateTodoList).Methods("PUT")
	mux.HandleFunc("/todo-lists/{id}", todoController.deleteTodoList).Methods("DELETE")

	return todoController
}

func (tl *TodoController) getTodoList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	todoPresenter := presenter.TodoPresenter{}
	result, err := tl.todoRepo.GetLists()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	todoPresenter.SendListResponse(result).JSON(w)
}

func (tl *TodoController) getTodoListDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	todoPresenter := presenter.TodoPresenter{}

	result, err := tl.todoRepo.GetOneById(id)
	if err != nil {
		fmt.Println(err.Error())
		todoPresenter.SendNotFoundResponse(err.Error()).JSON(w)
		return
	}

	todoPresenter.SendDetailResponse(result).JSON(w)
}

func (tl *TodoController) postTodoList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	todoPresenter := presenter.TodoPresenter{}
	reqBody := json.NewDecoder(r.Body)
	var todo model.Todo
	//pointing data json from body to _todo
	err := reqBody.Decode(&todo)
	if err != nil {
		fmt.Println(err.Error())
		todoPresenter.SendServerErrorResponse(err.Error()).JSON(w)
		return
	}
	fmt.Println(todo)
	tl.todoRepo.Post(todo)

	todoPresenter.SendSuccessResponse().JSON(w)
}

func (tl *TodoController) updateTodoList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	todoPresenter := presenter.TodoPresenter{}
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	reqBody := json.NewDecoder(r.Body)
	var todo model.Todo
	err := reqBody.Decode(&todo)
	if err != nil {
		fmt.Println(err.Error())
		todoPresenter.SendServerErrorResponse(err.Error()).JSON(w)
		return
	}
	tl.todoRepo.Update(todo,id)

	todoPresenter.SendSuccessResponse().JSON(w)
}

func (tl *TodoController) deleteTodoList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	todoPresenter := presenter.TodoPresenter{}
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	tl.todoRepo.Delete(id)

	todoPresenter.SendSuccessResponse().JSON(w)
}