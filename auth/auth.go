package auth

import (
	"encoding/base64"
	"errors"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewJwtService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	SECRET_KEY = os.Getenv("SECRET_KEY")
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	hashJwt := base64.StdEncoding.EncodeToString([]byte(SECRET_KEY))
	signedToken, err := token.SignedString([]byte(hashJwt))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, isValidAlogrithm := token.Method.(*jwt.SigningMethodHMAC)

		if !isValidAlogrithm {
			return nil, errors.New("Invalid Token")
		}

		hashJwt := base64.StdEncoding.EncodeToString([]byte(SECRET_KEY))
		return []byte(hashJwt), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
