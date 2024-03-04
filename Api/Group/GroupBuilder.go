package group

import (
	"github.com/gofiber/fiber/v2"
	middleware "go-api-test.kayn.ooo/Api/Middleware"
)

func AdminGroup(r fiber.Router) fiber.Router {
	return r.Group(
		"",
		middleware.IsLoggedIn,
		middleware.IsGranted([]string{"ROLE_ADMIN"}),
	)
}

func UserGroup(r fiber.Router) fiber.Router {
	return r.Group(
		"",
		middleware.IsLoggedIn,
		middleware.IsGranted([]string{"ROLE_USER"}),
	)
}
