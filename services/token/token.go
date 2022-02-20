package token

import (
	"cc_eduardherrera_BackendAPI/domain"
	"cc_eduardherrera_BackendAPI/entity"
	"cc_eduardherrera_BackendAPI/tools"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-pg/pg/v10"
)

type CustomClaims struct {
	Email string
	jwt.StandardClaims
}

func getSecret() interface{} {
	secretKey := tools.GetDotEnvVariable("SECRET_TOKEN_KEY", "029b937dab2b4e79e24757d5c316b785397b30b6938c71f7ff6d4e7665d0a046")
	return []byte(secretKey)
}

func GetToken(request entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email": request.Email,
	})
	signedToken, err := token.SignedString(getSecret())
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParseToken(receivedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return getSecret(), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func GetEmailFromToken(receivedToken string) (string, error) {
	parsedToken, err := ParseToken(receivedToken)
	var parsedTokenEmail string
	if err != nil {
		return "", err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		parsedTokenEmail = fmt.Sprint(claims["Email"])
	}

	return parsedTokenEmail, nil
}

func IsValidToken(db *pg.DB, token string) bool {
	parsedToken, err := ParseToken(token)
	var parsedTokenEmail string
	if err != nil {
		return false
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		parsedTokenEmail = fmt.Sprint(claims["Email"])
	}

	user, err := domain.FetchByEmail(db, parsedTokenEmail)
	if err != nil {
		return false
	}

	return user.Email != ""
}
