package token

import (
	"fmt"
	"net/smtp"
	"time"
	"travel/config"
	"travel/models"

	"github.com/golang-jwt/jwt"
)

func GenarateJWTToken(user *models.UserForLogin) (*models.ResponseRefreshToken, error) {
	accesToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	claimsAccess := accesToken.Claims.(jwt.MapClaims)
	claimsAccess["user_id"] = user.Id
	claimsAccess["username"] = user.Username
	claimsAccess["email"] = user.Email
	claimsAccess["full_name"] = user.FullName
	claimsAccess["iat"] = time.Now().Unix()
	claimsAccess["exp"] = time.Now().Add(time.Hour).Unix()

	access, err := accesToken.SignedString([]byte(config.Load().SINGNING_KEY_ACCESS))
	if err != nil {
		return nil, fmt.Errorf("error with generating access token: %s", err)
	}

	claimsRefresh := refreshToken.Claims.(jwt.MapClaims)
	claimsRefresh["user_id"] = user.Id
	claimsRefresh["username"] = user.Username
	claimsRefresh["email"] = user.Email
	claimsRefresh["full_name"] = user.FullName
	claimsRefresh["iat"] = time.Now().Unix()
	claimsRefresh["exp"] = time.Now().Add(time.Hour * 24).Unix()

	refresh, err := accesToken.SignedString([]byte(config.Load().SINGNING_KEY_REFRESH))
	if err != nil {
		return nil, fmt.Errorf("error with generating refresh token: %s", err)
	}

	return &models.ResponseRefreshToken{
		AccessToken:  access,
		RefreshToken: refresh,
		ExpiresIn:    int(time.Now().Add(time.Hour).Unix()),
	}, nil
}

func GenarateAccessToken(refreshToken string) (
	*models.ResponseRefreshToken, error) {
	accesToken := jwt.New(jwt.SigningMethodHS256)

	claims, err := ExtractClaims(refreshToken)
	if err != nil {
		return nil, err
	}
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	accesToken.Claims = claims
	access, err := accesToken.SignedString([]byte(config.Load().SINGNING_KEY_ACCESS))
	return &models.ResponseRefreshToken{
		AccessToken:  access,
		RefreshToken: refreshToken,
		ExpiresIn:    int(time.Now().Add(time.Hour).Unix()),
	}, err
}

func ValidateToken(tokenStr string) (bool, error) {
	_, err := ExtractClaims(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(
		t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v",
				t.Header["alg"])
		}
		return []byte(config.Load().SINGNING_KEY_REFRESH), nil
	})
	fmt.Println(config.Load().SINGNING_KEY_REFRESH)

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}

func SendEmail(to, subject, body string) error {
	// Set up authentication information.
	from := config.Load().EMAIL
	password := config.Load().PASSWORD
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Set up email content.
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body + "\r\n")

	// Send the email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
