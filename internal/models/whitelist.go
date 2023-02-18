package models

type Whitelist struct {
	BaseModel
	OperationName string
}

func (Whitelist) TableName() string {
	return "whitelists"
}
