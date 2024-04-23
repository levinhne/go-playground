package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Todo struct {
	bun.BaseModel `bun:"table:todos,alias:u"`

	ID    uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	Title string

	CreatedAt time.Time
	UpdatedAt time.Time
}
