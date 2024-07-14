package api

import (
	"database/sql"
	"travel/api/handler"

	_ "travel/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFile "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

// @title Auth Service
// @version 1.0
// @description This is the auth service of TravelTales app

// @contact.name Saidakbar
// @contact.url http://www.support_me_with_smile
// @contact.email "pardaboyevsaidakbar103@gmail.com"

// @host localhost:1111
// @BasePath /auth
func NewRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.GET("swagger/*any", swagger.WrapHandler(swaggerFile.Handler))

	auth := r.Group("/auth")
	h := handler.NewHandler(db)

	auth.POST("/register", h.Register)
	auth.POST("/login", h.Login)
	auth.POST("/reset-password", h.ResetPassword)
	auth.POST("/reset-password-part2/:email", h.ResetPasswordPart2)
	auth.POST("/refresh-token", h.RefreshToken)
	auth.POST("/logout", h.Logout)

	return r
}
