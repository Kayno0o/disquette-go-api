package entity

import (
	"github.com/uptrace/bun"
	trait "go-api-test.kayn.ooo/Api/Entity/Trait"
)

type Vote struct {
	bun.BaseModel `bun:"table:vote,alias:v"`
	trait.Identifier
	trait.Timestampable

	Up bool `bun:",notnull" json:"up"`

	AuthorId    uint       `bun:",notnull,pk,unique:author_disquette" json:"author_id"`
	DisquetteId uint       `bun:",notnull,pk,unique:author_disquette" json:"disquette_id"`
	Author      *User      `bun:"rel:belongs-to,join:author_id=id" json:"author"`
	Disquette   *Disquette `bun:"rel:belongs-to,join:disquette_id=id" json:"disquette"`
}
