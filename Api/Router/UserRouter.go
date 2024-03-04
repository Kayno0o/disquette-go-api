package router

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	entity "go-api-test.kayn.ooo/Api/Entity"
	fixture "go-api-test.kayn.ooo/Api/Fixture"
	group "go-api-test.kayn.ooo/Api/Group"
	repository "go-api-test.kayn.ooo/Api/Repository"
	security "go-api-test.kayn.ooo/Api/Security"
)

type UserRouter struct {
	GenericRouterInterface
}

func (ur *UserRouter) RegisterRoutes(r fiber.Router) {
	fixture.GenerateUsers(10, false)
	repository.UserRepository.Create(&entity.User{
		Username: "admin",
		Email:    "kevyn.fyleyssant@gmail.com",
		Password: security.HashPassword("password"),
		Roles:    []string{"ROLE_ADMIN"},
	})

	r.Post(
		"/user/login",
		ur.Login,
	).Post(
		"/user/register",
		ur.Register,
	)

	// ADMIN
	admin := group.AdminGroup(r)
	admin.Get(
		"/users/fixture/:amount",
		ur.Fixture,
	)

	// USER
	user := group.UserGroup(r)
	user.Get(
		"/user/me",
		ur.Me,
	)

	// PUBLIC
	r.Get(
		"/users",
		FindAll(
			repository.UserRepository,
			&[]entity.User{},
			&[]entity.UserContext{},
		),
	).Get(
		"/users/count",
		CountAll(
			repository.UserRepository,
			&entity.User{},
		),
	).Get(
		"/user/:id",
		FindOne(
			repository.UserRepository,
			&entity.User{},
			&entity.UserContext{},
		),
	)
}

func (r *UserRouter) Login(c *fiber.Ctx) error {
	var login entity.Login
	if err := c.BodyParser(&login); err != nil {
		return c.SendStatus(400)
	}

	user, err := security.Authenticate(&login)
	if err != nil {
		return c.Status(401).SendString("Unauthorized - login")
	}

	token, err := security.GenerateToken(user)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(token)
}

func (r *UserRouter) Register(c *fiber.Ctx) error {
	var form entity.Register
	if err := c.BodyParser(&form); err != nil {
		return c.SendStatus(400)
	}

	var user entity.User
	user.Username = form.Username
	user.Email = form.Email

	password := security.HashPassword(form.Password)
	user.Password = password

	_, err := repository.UserRepository.Create(&user)
	if err != nil {
		return c.SendStatus(500)
	}

	token, err := security.GenerateToken(&user)
	if err != nil {
		return c.SendStatus(500)
	}

	// add token to session/cookies
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token.Token,
		Path:     "/",
		Expires:  token.ExpiresAt,
		HTTPOnly: true,
		Domain:   "disquette.kayn.ooo",
		Secure:   true,
	})

	return c.JSON(token)
}

func (r *UserRouter) Fixture(c *fiber.Ctx) error {
	amount, err := strconv.Atoi(c.Params("amount"))
	if err != nil {
		return c.SendStatus(400)
	}

	users := fixture.GenerateUsers(amount, false)

	return c.JSON(users)
}

func (r *UserRouter) Me(c *fiber.Ctx) error {
	user := c.Locals("user")
	if user == nil {
		return c.Status(401).SendString("Unauthorized - me")
	}

	return c.JSON(user)
}

func (r *UserRouter) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Domain:   "disquette.kayn.ooo",
		Secure:   true,
	})

	return c.SendStatus(200)
}
