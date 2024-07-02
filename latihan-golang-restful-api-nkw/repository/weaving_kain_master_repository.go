package repository

import (
	"context"
	"database/sql"
	"latihan-golang-restful-api-nkw/model"
)

type WeavingKainMasterRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []model.WeavingKainMaster
}
