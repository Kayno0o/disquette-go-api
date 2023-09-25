package entity

import (
	"github.com/uptrace/bun"
	trait "go-api-test.kayn.ooo/Api/Entity/Trait"
)

type Disquette struct {
	bun.BaseModel `bun:"table:disquette,alias:d"`
	trait.Identifier
	trait.Timestampable

	AuthorId uint  `bun:",notnull" json:"author_id"`
	Author   *User `bun:"rel:belongs-to,join:author_id=id" json:"author"`

	IsOc    bool   `bun:",notnull" json:"is_oc"`
	Content string `bun:",notnull" json:"content"`

	Tags []Tag `bun:"m2m=disquette_tag,join:Disquette=Tag" json:"tags"`
}
