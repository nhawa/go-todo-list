package main

import (
	"go-todo-list/application"
)

func main() {
	apps := application.New()
	apps.ListenAndServe()
}
