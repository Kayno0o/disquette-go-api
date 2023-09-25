package trait

import "time"

type Timestampable struct {
	CreatedAt time.Time `bun:",nullzero,default:now()" json:"created_at"`
}
