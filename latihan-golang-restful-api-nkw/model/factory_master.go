package model

import "database/sql"

type FactoryMaster struct {
	Id           int64
	Name         string
	Type         string
	Status       bool
	CreatedDate  sql.NullTime
	ModifiedDate sql.NullTime
	CreatedBy    sql.NullInt64
	ModifiedBy   sql.NullInt64
}
