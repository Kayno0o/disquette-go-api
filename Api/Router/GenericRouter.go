package router

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	trait "go-api-test.kayn.ooo/Api/Entity/Trait"
	middleware "go-api-test.kayn.ooo/Api/Middleware"
	repository "go-api-test.kayn.ooo/Api/Repository"
)

var (
	UserRouter = UserRouterInterface{}
)

type GenericRouterInterface struct {
}

func Init() {
	fiberApp := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	api := fiberApp.Group("/api", middleware.Authenticate)
	UserRouter.RegisterUserRoutes(api)

	log.Fatal(fiberApp.Listen(":3000"))
}

func queryToParams(c *fiber.Ctx) map[string]interface{} {
	params := map[string]interface{}{}
	for key, value := range c.Queries() {
		params[key] = value
	}
	return params
}

func FindOne(rep repository.GenericRepositoryInterface, entity trait.EntityInterface) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(400)
		}
		rep.FindOneById(entity, id)
		if entity.GetId() == 0 {
			return c.SendStatus(404)
		}
		return c.JSON(entity)
	}
}

func FindAll(rep repository.GenericRepositoryInterface, entities interface{}) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		params := queryToParams(c)
		if params["offset"] == nil {
			params["offset"] = 1
		}
		if params["limit"] == nil {
			params["limit"] = 10
		}

		rep.FindAllBy(entities, params)
		return c.JSON(entities)
	}
}

func CountAll(rep repository.GenericRepositoryInterface, entity trait.EntityInterface) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		count, err := rep.CountAll(entity)
		if err != nil {
			return c.SendStatus(500)
		}
		return c.JSON(count)
	}
}
