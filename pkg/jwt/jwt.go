package jwt

import (
	"Gin/models"
	"Gin/pkg/initconf"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

var jwtSecret = []byte(initconf.JwtSecret)

// Cteate token.
func CreateToken(username, password string) (string, error) {
	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: int64(time.Now().Unix() + 1000),
		Issuer:    "Capwang",
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// secret add in token
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		log.Fatalf("[jwt] failed to add secret into token : %v\n", err)
	}
	return token, err
}

// parse token
func ParseToken(token string) (*models.Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		log.Fatalf("[jwt] failed to parse token : %v\n", err)
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*models.Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
