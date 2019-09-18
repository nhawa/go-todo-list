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
		todoRepo         *repository.TodoRepository
	}
)


func NewTodoController(
	mux *mux.Router,
	todoRepo *repository.TodoRepository,
) *TodoController {

	todoController := &TodoController{}
	todoController.todoRepo = todoRepo

	getTodoListHandler := http.HandlerFunc(todoController.getTodoList)
	postTodoListHandler := http.HandlerFunc(todoController.postTodoList)
	getTodoListDetailHandler := http.HandlerFunc(todoController.getTodoListDetail)
	updateTodoListHandler := http.HandlerFunc(todoController.updateTodoList)
	deleteTodoListHandler := http.HandlerFunc(todoController.deleteTodoList)

	mux.Handle("/", middleware.JWT(getTodoListHandler)).Methods("GET")
	mux.Handle("/todo-lists", middleware.JWT(getTodoListHandler)).Methods("GET")
	mux.Handle("/todo-lists", middleware.JWT(postTodoListHandler)).Methods("POST")
	mux.Handle("/todo-lists/{id}", middleware.JWT(getTodoListDetailHandler)).Methods("GET")
	mux.Handle("/todo-lists/{id}", middleware.JWT(updateTodoListHandler)).Methods("PUT")
	mux.Handle("/todo-lists/{id}", middleware.JWT(deleteTodoListHandler)).Methods("DELETE")

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

	todoPresenter.SendSuccessUpdatedResponse().JSON(w)
}

func (tl *TodoController) deleteTodoList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	todoPresenter := presenter.TodoPresenter{}
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	tl.todoRepo.Delete(id)

	todoPresenter.SendSuccessDeletedResponse().JSON(w)
}