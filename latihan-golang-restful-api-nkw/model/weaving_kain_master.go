package model

import "database/sql"

type WeavingKainMaster struct {
	Id           int64
	KainName     string
	Anyaman      string
	ActiveStatus bool
	LebarKain    float64
	PanjangKain  float64
	UpdateDate   sql.NullTime
	Status       bool
	CreatedDate  sql.NullTime
	ModifiedDate sql.NullTime
	CreatedBy    sql.NullInt64
	ModifiedBy   sql.NullInt64
}
