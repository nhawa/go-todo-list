package repository

import (
	"database/sql"
	"fmt"
	"go-todo-list/authentication/model"
)

type(
	UserRepository struct {
		DB *sql.DB
	}
)

func NewUserRepository(DB *sql.DB) *UserRepository {
	user := &UserRepository{}
	user.DB = DB

	return user
}


func (userRepository *UserRepository) FindOneForCustomerLogin(email string) (*model.User, error) {
	user := new(model.User)
	query := `
		select
			u.id,
			u.email,
			u.salt,
			u.password,
			u.name
		from
			user_account u
-- 		inner join
-- 			customer c
-- 		on
-- 			c.user_account_id=u.id
		where
			email=$1 limit 1
	`

	err := userRepository.DB.QueryRow(query, email).
		Scan(&user.ID, &user.Email, &user.Salt, &user.Password, &user.Name)

	if err != nil {
		fmt.Println("error get user")
		return nil, err
	}

	return user, nil
}