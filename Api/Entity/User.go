package entity

import (
	"github.com/uptrace/bun"
	trait "go-api-test.kayn.ooo/Api/Entity/Trait"
)

type User struct {
	bun.BaseModel `bun:"table:user,alias:u"`
	trait.Identifier
	trait.Timestampable

	Username string   `bun:",notnull" json:"username"`
	Email    string   `bun:",notnull,unique" json:"email"`
	Password string   `bun:",notnull" json:"password"`
	Roles    []string `bun:",array" json:"roles"`
}

type UserContext struct {
	Username string `json:"username"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) HasRole(role string) bool {
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}

	return false
}
