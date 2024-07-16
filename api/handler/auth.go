package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"time"
	"travel/api/token"
	"travel/models"

	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
)

// @Summery Register User
// @Description this API for register new user to the TravelTales app
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body models.RequestRegister true "All params are required"
// @Success 200 {object} models.ResponseRegister "Returns new User's information"
// @Failure 400 {object} models.Error "it occur when user send invalid parametrs"
// @Failure 500 {object} models.Error "it occur when error with write user params into database"
// @Router /register [post]
func (h *Handler) Register(ctx *gin.Context) {
	var req models.RequestRegister

	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error with decoding url body: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": fmt.Sprintf("Error with decoding url body: %s", err.Error()),
		})
		return
	}

	// check email is valid
	if _, err := mail.ParseAddress(req.Email); err != nil {
		h.Logger.Error(fmt.Sprintf("invalid email: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": fmt.Sprintf("invalid email: %s", err.Error()),
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error with hashing password: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with hashing password: %s", err.Error()),
		})
		return
	}
	req.Password = string(hashedPassword)

	resp, err := h.AuthRepo.Register(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error with register User: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with register User: %s", err.Error()),
		})
		return
	}
	fmt.Println(resp)
	ctx.JSON(http.StatusOK, resp)
}

// @Summery Login User
// @Description this API for login new user to the TravelTales app
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body models.RequestLogin true "email and password required"
// @Success 200 {object} models.ResponseRefreshToken "Returns access and refresh tokens"
// @Failure 400 {object} models.Error "it occur when user send invalid email or password"
// @Failure 500 {object} models.Error "it occur when error with write user params into database"
// @Router /login [post]
func (h *Handler) Login(ctx *gin.Context) {
	req := models.RequestLogin{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error with decoding url body: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": fmt.Sprintf("Error with decoding url body: %s", err.Error()),
		})
		return
	}

	user, err := h.AuthRepo.GetUserByEmail(req.Email)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("error getting email from database %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": fmt.Sprintf("error getting email from database %s", err.Error()),
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Password is incorrect %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": fmt.Sprintf("Password is incorrect %s", err.Error()),
		})
		return
	}

	tokens, err := token.GenarateJWTToken(user)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error while generating tokens %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error while generating tokens %s", err.Error()),
		})
		return
	}

	err = h.AuthRepo.DeleteRefreshTokenByUserId(user.Id)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error with deleting userinfo from refreshToken table %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with deleting userinfo from refreshToken table %s", err.Error()),
		})
		return
	}

	err = h.AuthRepo.StoreRefreshToken(&models.StoreRefreshToken{
		UserId:       user.Id,
		RefreshToken: tokens.RefreshToken,
		ExpiresIn:    time.Now().Add(time.Hour * 24),
	})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error with deleting userinfo from refreshToken table %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with deleting userinfo from refreshToken table %s", err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, tokens)
}

// @Summery reset user password
// @Description this API for reset user password to the TravelTales app
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body models.RequestResetPassword true "email is required"
// @Success 200 {object} models.ResponseResetPassword "Returns message about sending message to email"
// @Failure 400 {object} models.Error "it occur when user send invalid email token"
// @Failure 500 {object} models.Error "it occur when error with sending message to the email"
// @Router /reset-password [post]
func (h *Handler) ResetPassword(ctx *gin.Context) {

	req := models.RequestResetPassword{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error with decoding url body: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": fmt.Sprintf("Error with decoding url body: %s", err.Error()),
		})
		return
	}

	// check email is valid
	if _, err := mail.ParseAddress(req.Email); err != nil {
		h.Logger.Error(fmt.Sprintf("invalid email: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": fmt.Sprintf("invalid email: %s", err.Error()),
		})
		return
	}

	exists, err := h.AuthRepo.EmailExists(req.Email)
	if !exists || err != nil {
		h.Logger.Error(fmt.Sprintf("email is not registered: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": fmt.Sprintf("email is not registered: %s", err.Error()),
		})
		return
	}

	subject := "reset-password of TravelTales app"
	message := "CLick here to reset password\n\nhttp://localhost:1111/swagger/users/index.html#/Auth/post_reset_password_part2/"

	err = token.SendEmail(req.Email, subject, message+req.Email)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error with decoding url body: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": fmt.Sprintf("Error with decoding url body: %s", err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, models.ResponseResetPassword{
		Message: "sent message to the email",
	})
}

// @Summery reset user password part2
// @Description this API for user who receive message from his/her email to the TravelTales app
// @Tags Auth
// @Accept json
// @Produce json
// @Param email path string true "email is required"
// @Param request body models.UpdatePassword true "enter new password"
// @Success 200 {object} models.ResponseResetPassword "Returns message about reset-password"
// @Failure 400 {object} models.Error "it occur when user enter invalid password"
// @Failure 500 {object} models.Error "it occur when error with reseting pasword"
// @Router /reset-password-part2/{email} [post]
func (h *Handler) ResetPasswordPart2(ctx *gin.Context) {

	req := models.UpdatePassword{}
	email := ctx.Param("email")

	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil || email == "" {
		h.Logger.Error(fmt.Sprintf("Error with taking email and password from frontend: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": fmt.Sprintf("Error with taking email and password from frontend: %s", err.Error()),
		})
		return
	}

	exists, err := h.AuthRepo.EmailExists(email)
	if !exists || err != nil {
		h.Logger.Error(fmt.Sprintf("email is not registered: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": fmt.Sprintf("email is not registered: %s", err.Error()),
		})
		return
	}

	err = h.AuthRepo.UpdatePassword(email, req.Password)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error with writing password to db: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": fmt.Sprintf("Error with writing password to db: %s", err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, models.ResponseResetPassword{
		Message: "Password successfully updated",
	})
}

// @Summery refresh user access token
// @Description this API for refresh user access token to the TravelTales app
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body models.RequestRefreshToken true "refresh token is required"
// @Success 200 {object} models.ResponseRefreshToken "Returns access and refresh tokens"
// @Failure 400 {object} models.Error "it occur when user send invalid refresh token"
// @Failure 500 {object} models.Error "it occur when error with generating access token"
// @Router /refresh-token [post]
func (h *Handler) RefreshToken(ctx *gin.Context) {
	req := models.RequestRefreshToken{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error with decoding url body: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": fmt.Sprintf("Error with decoding url body: %s", err.Error()),
		})
		return
	}

	check, err := h.AuthRepo.CheckRefreshTokenExists(req.RefreshToken)
	if err != nil || !check {
		h.Logger.Error(fmt.Sprintf("Invalid refresh token: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": fmt.Sprintf("Invalid refresh token: %s", err.Error()),
		})
		return
	}

	resp, err := token.GenarateAccessToken(req.RefreshToken)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error with generate access token: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with generate access token: %s", err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summery logout user
// @Description this API for logout user  to the TravelTales app
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body models.RequestLogout true "refresh token is required"
// @Success 200 {object} models.ResponseLogout "returns massage for success"
// @Failure 400 {object} models.Error "it occur when user send invalid refresh token"
// @Failure 500 {object} models.Error "it occur when error with error of deleting refresh token"
// @Router /logout [post]
func (h *Handler) Logout(ctx *gin.Context) {
	req := models.RequestLogout{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error with decoding url body: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": fmt.Sprintf("Error with decoding url body: %s", err.Error()),
		})
		return
	}

	if err = h.AuthRepo.DeleteRefreshToken(req.RefreshToken); err != nil {
		h.Logger.Error(fmt.Sprintf("Error with Delete refreshtoken from database: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with Delete refreshtoken from database: %s", err.Error()),
		})
		return
	}

	resp := models.ResponseLogout{
		Message: "Successfully logged out",
	}
	ctx.JSON(http.StatusOK, resp)
}
