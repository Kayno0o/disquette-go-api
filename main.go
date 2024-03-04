package main

import (
	"github.com/joho/godotenv"
	repository "go-api-test.kayn.ooo/Api/Repository"
	router "go-api-test.kayn.ooo/Api/Router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	rep := repository.GenericRepository{}

	rep.Init()
	router.Init([]router.GenericRouterInterface{
		&router.UserRouter{},
	})
}
