package router

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	trait "go-api-test.kayn.ooo/Api/Entity/Trait"
	middleware "go-api-test.kayn.ooo/Api/Middleware"
	repository "go-api-test.kayn.ooo/Api/Repository"
	utils "go-api-test.kayn.ooo/Api/Utils"
)

type GenericRouterInterface interface {
	RegisterRoutes(fiber.Router)
}

func Init(routers []GenericRouterInterface) {
	fiberApp := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	api := fiberApp.Group("/api", middleware.Authenticate)
	for i := range routers {
		routers[i].RegisterRoutes(api)
	}

	log.Fatal(fiberApp.Listen(":3000"))
}

func queryToParams(c *fiber.Ctx) map[string]interface{} {
	params := map[string]interface{}{}
	for key, value := range c.Queries() {
		params[key] = value
	}
	return params
}

func FindOne(rep repository.GenericRepositoryInterface, entity trait.IdentifierInterface, context interface{}) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(400)
		}
		rep.FindOneById(entity, id)
		if entity.GetId() == 0 {
			return c.SendStatus(404)
		}
		utils.ApplyContext(entity, context)
		return c.JSON(context)
	}
}

func FindAll(rep repository.GenericRepositoryInterface, entities interface{}, context interface{}) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		params := queryToParams(c)
		if params["offset"] == nil {
			params["offset"] = 1
		}
		if params["limit"] == nil {
			params["limit"] = 10
		}

		rep.FindAllBy(entities, params)
		utils.ApplyContext(entities, context)
		return c.JSON(context)
	}
}

func CountAll(rep repository.GenericRepositoryInterface, entity trait.IdentifierInterface) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		count, err := rep.CountAll(entity)
		if err != nil {
			return c.SendStatus(500)
		}
		return c.JSON(count)
	}
}
