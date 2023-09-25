package entity

import (
	"github.com/uptrace/bun"
	trait "go-api-test.kayn.ooo/Api/Entity/Trait"
)

type DisquetteTag struct {
	bun.BaseModel `bun:"table:disquette_tag,alias:dt"`
	trait.Identifier
	trait.Timestampable

	DisquetteId uint       `bun:",notnull,pk,unique:disquette_tag" json:"disquette_id"`
	TagId       uint       `bun:",notnull,pk,unique:disquette_tag" json:"tag_id"`
	Disquette   *Disquette `bun:"rel:belongs-to,join:disquette_id=id" json:"disquette"`
	Tag         *Tag       `bun:"rel:belongs-to,join:tag_id=id" json:"tag"`
}
