package models

type Whitelist struct {
	BaseModel
	QueryID int `db:"query_id"`
}

func (Whitelist) TableName() string {
	return "whitelists"
}
