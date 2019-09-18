package repository

import (
	"database/sql"
	"time"
)

type (
	AccessTokenRepositoryInterface interface {
		Save(token string, userId int, expiration time.Time, refreshToken string) error
	}
	accessTokenRepository struct {
		DB *sql.DB
	}
)

func NewTokenRepository(sql *sql.DB) AccessTokenRepositoryInterface {
	return &accessTokenRepository{
		DB: sql,
	}
}

func (a *accessTokenRepository) Save(token string, userID int, expiration time.Time, refreshToken string) error {

	query := `
		insert into access_token (
			token,
			otten_account_id,
			expired_at,
			refresh_token
		) values (
			$1,
			$2,
			$3,
			$4
		)
	`

	_, err := a.DB.Exec(query, token, userID, expiration, refreshToken)

	if err != nil {
		return err
	}

	return nil
}
