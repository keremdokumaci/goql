package models

import "time"

type QueryCache struct {
	BaseModel
	QueryID int       `db:"query_id"`
	Ttl     time.Time `db:"ttl"`
}

func (QueryCache) TableName() string {
	return "query_caches"
}
