package model

import "time"

type Todo struct {
	Id          int    		`db:"id"`
	Name        string 		`db:"name"`
	Description string 		`db:"description"`
	DueDate     time.Time 	`db:"due_date"`
	Status      string 		`db:"status"`
	CreatedAt   time.Time 	`db:"created_at"`
}