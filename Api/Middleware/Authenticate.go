package middleware

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	entity "go-api-test.kayn.ooo/Api/Entity"
	repository "go-api-test.kayn.ooo/Api/Repository"
)

func Authenticate(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		tokenString = c.Cookies("token")
	}

	if tokenString == "" {
		return c.Next()
	}

	if tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	}, jwt.WithValidMethods([]string{"HS256"}), jwt.WithAudience("disquette.kayn.ooo"), jwt.WithIssuer("disquette.kayn.ooo"))

	if err != nil || !token.Valid {
		c.Status(401).SendString("Invalid token")
		return nil
	}

	claims := token.Claims.(jwt.MapClaims)

	if claims["exp"] == nil || claims["iat"] == nil {
		c.Status(401).SendString("Invalid token")
		return nil
	}

	if int64(claims["exp"].(float64)) < time.Now().Unix() || int64(claims["iat"].(float64)) > time.Now().Unix() {
		c.Status(401).SendString("Invalid token")
		return nil
	}

	id := claims["id"].(float64)

	user := &entity.User{}
	repository.UserRepository.FindOneById(user, int(id))
	c.Locals("user", user)

	return c.Next()
}
