package fixture

import (
	"math/rand"
	"os"
	"strings"

	entity "go-api-test.kayn.ooo/Api/Entity"
	repository "go-api-test.kayn.ooo/Api/Repository"
	security "go-api-test.kayn.ooo/Api/Security"
)

func GetFirstNames() []string {
	file, err := os.ReadFile("./firstnames.txt")
	if err != nil {
		panic(err)
	}

	return strings.Split(string(file), "\n")
}

func RandomFirstName(firstNames []string) string {
	return firstNames[rand.Intn(len(firstNames))]
}

func GenerateUsers(nb int, isAdmin bool) []entity.User {
	users := []entity.User{}
	firstNames := GetFirstNames()
	for i := 0; i < nb; i++ {
		firstName := RandomFirstName(firstNames)

		password := security.HashPassword("password")

		role := "ROLE_USER"
		if isAdmin {
			role = "ROLE_ADMIN"
		}

		user := &entity.User{
			Username: firstName,
			Email:    firstName + "@gmail.com",
			Password: password,
			Roles:    []string{role},
		}

		users = append(users, *user)
	}

	_, err := repository.DB.NewInsert().Model(&users).Exec(repository.Ctx)
	if err != nil {
		panic(err)
	}

	return users
}
