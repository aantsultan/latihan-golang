package repository

import (
	"context"
	"database/sql"
	"latihan-golang-restful-api-nkw/model"
)

type FactoryMasterRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []model.FactoryMaster
}
