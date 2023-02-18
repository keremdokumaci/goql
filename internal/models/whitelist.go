package models

type Whitelist struct {
	BaseModel
	OperationName string `db:"operation_name"`
}

func (Whitelist) TableName() string {
	return "whitelists"
}
