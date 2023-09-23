package security

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	entity "go-api-test.kayn.ooo/Api/Entity"
	repository "go-api-test.kayn.ooo/Api/Repository"
	"golang.org/x/crypto/bcrypt"
)

type JWT struct {
	Token string `json:"token"`
}

type Claims struct {
	Id uint `json:"id"`
	jwt.Claims
}

func GenerateToken(user *entity.User) (*JWT, error) {
	// Load the secret key from the .env file
	secretKey := []byte(os.Getenv("SECRET_KEY"))

	// Create a new JWT token with the user ID and expiration time
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["aud"] = "disquette.kayn.ooo"
	claims["iss"] = "disquette.kayn.ooo"

	// Sign the token with the secret key and return a JWT struct with the token string
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}
	return &JWT{Token: tokenString}, nil
}

func Authenticate(login *entity.Login) (*entity.User, error) {
	user := entity.User{}

	err := repository.UserRepository.FindOneBy(&user, map[string]interface{}{
		"email": login.Email,
	})
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, err
	}

	return &user, nil
}
