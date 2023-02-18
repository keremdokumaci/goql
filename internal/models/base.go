package models

import "time"

type Modeler interface {
	TableName() string
}

type BaseModel struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
}
