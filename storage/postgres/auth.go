package postgres

import (
	"database/sql"
	"time"
	"travel/models"

	"github.com/google/uuid"
)

type AuthRepo struct {
	DB *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{
		DB: db,
	}
}

func (a *AuthRepo) Register(req *models.RequestRegister) (
	*models.ResponseRegister, error) {

	query := `
		insert into users (
			id, username, email, password, full_name, created_at 
		) values (
			$1, $2, $3, $4, $5, $6
		)`

	newId := uuid.NewString()
	createdAt := time.Now()
	_, err := a.DB.Exec(query, newId, req.Username, req.Email, req.Password,
		req.FullName, createdAt)

	if err != nil {
		return nil, err
	}
	resp := models.ResponseRegister{
		Id:        newId,
		Username:  req.Username,
		Email:     req.Email,
		FullName:  req.FullName,
		CreatedAt: createdAt.Format(time.RFC3339),
	}
	return &resp, nil
}

func (a *AuthRepo) GetUserByEmail(email string) (
	*models.UserForLogin, error) {

	query := `
		select
			id, username, email, password, full_name, created_at
		from
			users
		where
			email = $1 and 
			deleted_at is null
	`
	user := models.UserForLogin{}
	err := a.DB.QueryRow(query, email).Scan(&user.Id, &user.Username,
		&user.Email, &user.Password, &user.FullName, &user.CreatedAt)

	return &user, err
}

func (a *AuthRepo) CheckRefreshTokenExists(refteshToken string) (bool, error) {

	query := `
		select
			1
		from 
			refresh_tokens
		where
			refresh_token = $1
	`

	res := 0
	err := a.DB.QueryRow(query, refteshToken).Scan(&res)
	if err != nil || err == sql.ErrNoRows {
		return false, err
	}
	return true, nil
}

func (a *AuthRepo) DeleteRefreshToken(refteshToken string) error {

	query := `
		delete from
			refresh_tokens
		where
			refresh_token = $1
	`

	_, err := a.DB.Exec(query, refteshToken)
	if err != nil || err == sql.ErrNoRows {
		return err
	}
	return nil
}

func (a *AuthRepo) StoreRefreshToken(refreshTokenInfo *models.StoreRefreshToken) error {

	query := `
		insert into refresh_tokens (
			id, user_id, refresh_token, expires_in
		) values (
			$1, $2, $3, $4 
		)
	`
	newId := uuid.NewString()
	_, err := a.DB.Exec(query, newId, refreshTokenInfo.UserId,
		refreshTokenInfo.RefreshToken, refreshTokenInfo.ExpiresIn)

	return err
}

func (a *AuthRepo) DeleteRefreshTokenByUserId(userid string) error {

	query := `
		delete from
			refresh_tokens
		where
			user_id = $1
	`

	_, err := a.DB.Exec(query, userid)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthRepo) EmailExists(email string) (bool, error) {

	query := `
		select
			1
		from 
			users
		where
			email = $1
	`

	res := 0
	err := a.DB.QueryRow(query, email).Scan(&res)
	if err != nil || err == sql.ErrNoRows {
		return false, err
	}
	return true, nil
}

func (a *AuthRepo) UpdatePassword(email, password string) error {

	query := `
		update
			users
		set
			password = $1
		where
			email = $2
	`

	_, err := a.DB.Exec(query, password, email)
	if err != nil {
		return err
	}
	return nil
}
