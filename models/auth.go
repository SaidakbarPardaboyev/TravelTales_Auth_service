package models

import "time"

type RequestRegister struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

type ResponseRegister struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FullName  string `json:"full_name"`
	CreatedAt string `json:"created_at"`
}

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseLogin struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

type RequestResetPassword struct {
	Email string `json:"email"`
}

type ResponseResetPassword struct {
	Message string `json:"message"`
}

type RequestRefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

type ResponseRefreshToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

type RequestLogout struct {
	RefreshToken string `json:"refresh_token"`
}

type ResponseLogout struct {
	Message string `json:"message"`
}

type Error struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type UserForLogin struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FullName  string `json:"full_name"`
	CreatedAt string `json:"created_at"`
}

type StoreRefreshToken struct {
	UserId       string    `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresIn    time.Time `json:"expires_in"`
}

type UpdatePassword struct {
	Password string `json:"password"`
}