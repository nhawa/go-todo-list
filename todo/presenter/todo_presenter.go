package presenter

import (
	"go-todo-list/lib"
	"go-todo-list/todo/model"
	"net/http"
)

const (
	GET_TODO_NOT_FOUND_MESSAGE  string = "Failed to fetch Todo List(s) data"
	GET_TODO_SERVER_ERROR_MESSAGE  string = "OooppS! something went wrong!!!"
	GET_TODO_SUCCESS_MESSAGE    string = "Todo list(s) data fetched successfully"
	GET_TODO_SUCCESS_UPDATE_MESSAGE    string = "Todo data updated successfully"
	GET_TODO_SUCCESS_DELETE_MESSAGE    string = "Todo data deleted successfully"

	GET_TODO_NOT_FOUND_CODE  int = 404
	GET_TODO_SERVER_ERROR_CODE  int = 500
	GET_TODO_SUCCESS_CODE  int = 200
)

type (
	TodoPresenter struct{}

	TodoListResponse struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}

	TodoDetailResponse struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		DueDate     string `json:"due_date"`
		Status      string `json:"status"`
		CreatedAt   string `json:"created_at"`
	}
)

func (tp *TodoPresenter) SendListResponse(todoLists []model.Todo) lib.Response {
	todoListResponse := make([]TodoListResponse, len(todoLists))

	for k, todo := range todoLists {

		todoList := TodoListResponse{
			Id:                todo.Id,
			Name:     		   todo.Name,
			Description:       todo.Description,
			Status:            todo.Status,
		}

		todoListResponse[k] = todoList
	}

	return lib.CreateResponse(
		http.StatusOK,
		GET_TODO_SUCCESS_CODE,
		GET_TODO_SUCCESS_MESSAGE,
		todoListResponse,
	)
}

func (tp *TodoPresenter) SendDetailResponse(todo model.Todo) lib.Response {

	todoDetailResponse := TodoDetailResponse{
		Id:                todo.Id,
		Name:     		   todo.Name,
		CreatedAt:         todo.CreatedAt.Format("2006-01-02 15:04:05"),
		DueDate:           todo.DueDate.Format("2006-01-02 15:04:05"),
		Description:       todo.Description,
		Status:            todo.Status,
	}

	return lib.CreateResponse(
		http.StatusOK,
		GET_TODO_SUCCESS_CODE,
		GET_TODO_SUCCESS_MESSAGE,
		todoDetailResponse,
	)
}

func (tp *TodoPresenter) SendNotFoundResponse(errorMessage string) lib.Response {
	return lib.CreateFailResponse(
		http.StatusNotFound,
		GET_TODO_NOT_FOUND_CODE,
		GET_TODO_NOT_FOUND_MESSAGE,
		nil,
		errorMessage,
	)
}

func (tp *TodoPresenter) SendServerErrorResponse(errorMessage string) lib.Response {
	return lib.CreateFailResponse(
		http.StatusNotFound,
		GET_TODO_SERVER_ERROR_CODE,
		GET_TODO_SERVER_ERROR_MESSAGE,
		nil,
		errorMessage,
	)
}

func (tp *TodoPresenter) SendSuccessResponse() lib.Response {
	return lib.CreateResponse(
		http.StatusNotFound,
		GET_TODO_SUCCESS_CODE,
		GET_TODO_SUCCESS_MESSAGE,
		nil,
	)
}

func (tp *TodoPresenter) SendSuccessDeletedResponse() lib.Response {
	return lib.CreateResponse(
		http.StatusNotFound,
		GET_TODO_SUCCESS_CODE,
		GET_TODO_SUCCESS_DELETE_MESSAGE,
		nil,
	)
}

func (tp *TodoPresenter) SendSuccessUpdatedResponse() lib.Response {
	return lib.CreateResponse(
		http.StatusNotFound,
		GET_TODO_SUCCESS_CODE,
		GET_TODO_SUCCESS_UPDATE_MESSAGE,
		nil,
	)
}