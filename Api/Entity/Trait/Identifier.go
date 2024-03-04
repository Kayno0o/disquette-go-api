package trait

type IdentifierInterface interface {
	GetId() int64
}

type Identifier struct {
	IdentifierInterface `bun:"-" json:"-"`

	ID int64 `bun:",pk,autoincrement" json:"id"`
}

func (i *Identifier) GetId() int64 {
	return i.ID
}
