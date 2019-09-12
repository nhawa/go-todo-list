package repository

import (
	"database/sql"
	"fmt"
	"go-todo-list/todo/model"
)

type (
	TodoRepository interface{
		GetLists() ([]model.Todo, error)
		GetOneById(id int) (model.Todo, error)
		Post(todo model.Todo)
		Update(todo model.Todo, id int)
		Delete(id int)
	}

	todoRepository struct {
		db *sql.DB
	}
)

func NewTodoRepository(DB *sql.DB) TodoRepository {
	repo := &todoRepository{db: DB}

	return repo
}

func (tr *todoRepository) GetLists() ([]model.Todo, error){
	var result []model.Todo

	rows, err := tr.db.Query("select * from todo")
	if err != nil {
		fmt.Println(err.Error())
		return result, err
	}
	defer rows.Close()


	for rows.Next() {
		var each = model.Todo{}
		var err = rows.Scan(&each.Id, &each.Name, &each.Description, &each.DueDate, &each.Status, &each.CreatedAt)

		if err != nil {
			fmt.Println(err.Error())
			return result, err
		}

		result = append(result, each)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return result, err
	}

	return result, nil
}

func (tr *todoRepository) GetOneById(id int) (model.Todo, error){
	var result = model.Todo{}

	err := tr.db.QueryRow("select * from todo where id = $1 limit 1", id).
		Scan(&result.Id, &result.Name, &result.Description, &result.DueDate, &result.Status, &result.CreatedAt)

	return result, err
}

func (tr *todoRepository) Post(todo model.Todo)  {
	_, err := tr.db.Query("insert into todo (name, description, due_date) values ($1, $2, $3)", todo.Name, todo.Description, todo.DueDate)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func (tr *todoRepository) Update(todo model.Todo, id int)  {
	_, err := tr.db.Query("update todo set name = $1, description = $2, due_date = $3 where id = $4", todo.Name, todo.Description, todo.DueDate, id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func (tr *todoRepository) Delete(id int)  {
	_, err := tr.db.Query("delete from todo where id = $1", id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}