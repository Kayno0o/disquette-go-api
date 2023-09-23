package router

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	entity "go-api-test.kayn.ooo/Api/Entity"
	fixture "go-api-test.kayn.ooo/Api/Fixture"
	middleware "go-api-test.kayn.ooo/Api/Middleware"
	repository "go-api-test.kayn.ooo/Api/Repository"
	security "go-api-test.kayn.ooo/Api/Security"
)

type UserRouterInterface struct {
}

func (ur *UserRouterInterface) RegisterUserRoutes(r fiber.Router) {
	fixture.GenerateUsers(10, false)
	repository.UserRepository.Create(&entity.User{
		Username: "admin",
		Email:    "kevyn.fyleyssant@gmail.com",
		Password: security.HashPassword("password"),
		Roles:    []string{"ROLE_ADMIN"},
	})

	r.Post("/user/login", ur.Login)
	r.Post("/user/register", ur.Register)

	// ADMIN
	admin := r.Group("", middleware.IsLoggedIn, middleware.IsGranted([]string{"ROLE_ADMIN"}))

	admin.Get("/users/fixture/:amount", ur.Fixture)

	// LOGGED
	logged := r.Group("", middleware.IsLoggedIn, middleware.IsGranted([]string{"ROLE_USER"}))

	logged.Get("/user/me", ur.Me)

	// PUBLIC
	r.Get("/users", FindAll(repository.UserRepository, &[]entity.User{}))
	r.Get("/users/count", CountAll(repository.UserRepository, &entity.User{}))
	r.Get("/user/:id", FindOne(repository.UserRepository, &entity.User{}))
}

func (r *UserRouterInterface) Login(c *fiber.Ctx) error {
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

func (r *UserRouterInterface) Register(c *fiber.Ctx) error {
	var user entity.User
	var form entity.Register
	if err := c.BodyParser(&form); err != nil {
		return c.SendStatus(400)
	}

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

	return c.JSON(token)
}

func (r *UserRouterInterface) Fixture(c *fiber.Ctx) error {
	amount, err := strconv.Atoi(c.Params("amount"))
	if err != nil {
		return c.SendStatus(400)
	}

	users := fixture.GenerateUsers(amount, false)

	return c.JSON(users)
}

func (r *UserRouterInterface) Me(c *fiber.Ctx) error {
	user := c.Locals("user")
	if user == nil {
		return c.Status(401).SendString("Unauthorized - me")
	}

	return c.JSON(user)
}
