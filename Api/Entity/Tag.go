package entity

import (
	"github.com/uptrace/bun"
	trait "go-api-test.kayn.ooo/Api/Entity/Trait"
)

type Tag struct {
	bun.BaseModel `bun:"table:tag,alias:t"`
	trait.Identifier
	trait.Timestampable

	Code  string `bun:",notnull" json:"code"`
	Label string `bun:",notnull" json:"label"`
}
