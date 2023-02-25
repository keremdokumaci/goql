package models

type Query struct {
	BaseModel
	Name string `db:"name"`
}

func (Query) TableName() string {
	return "queries"
}
