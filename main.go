package main

import (
	"github.com/joho/godotenv"
	entity "go-api-test.kayn.ooo/Api/Entity"
	repository "go-api-test.kayn.ooo/Api/Repository"
	router "go-api-test.kayn.ooo/Api/Router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	rep := repository.GenericRepository{}
	rep.Init([]interface{}{
		&entity.User{},
	})

	router.Init([]router.GenericRouterInterface{
		&router.UserRouter{},
	})
}
