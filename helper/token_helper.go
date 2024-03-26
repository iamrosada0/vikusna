package helper

import (
	"evaeats/domain"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtkey = []byte(os.Getenv("JWT_KEY"))

/*Genearte JWT token*/
func GenerateToken(principal domain.User, duration time.Duration) (string, int64, error) {

	claims := &domain.Claims{
		User_id:   principal.ID,
		User_name: principal.User_name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)

	if err != nil {
		return "", 0, err
	}

	return tokenString, claims.ExpiresAt, nil
}
