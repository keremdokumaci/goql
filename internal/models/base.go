package models

import (
	"database/sql"
	"time"
)

type Modeler interface {
	TableName() string
}

type BaseModel struct {
	ID        int          `db:"id"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
