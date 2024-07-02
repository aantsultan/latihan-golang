package repository

import (
	"context"
	"database/sql"
	"latihan-golang-restful-api-nkw/helper"
	"latihan-golang-restful-api-nkw/model"
)

type WeavingKainMasterRepositoryImpl struct {
}

func NewWeavingKainMasterRepository() *WeavingKainMasterRepositoryImpl {
	return &WeavingKainMasterRepositoryImpl{}
}

func (repository WeavingKainMasterRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []model.WeavingKainMaster {
	SQL := "select kain_id, kain_name, anyaman, active_status, lebar_kain, " +
		" panjang_kain, update_date, status, created_date, modified_date, " +
		" created_by, modified_by " +
		" from s_weaving_kain_master"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var weavingKainMasters []model.WeavingKainMaster
	if rows.Next() {
		kainMaster := model.WeavingKainMaster{}
		err := rows.Scan(
			&kainMaster.Id, &kainMaster.KainName, &kainMaster.Anyaman, &kainMaster.ActiveStatus, &kainMaster.LebarKain,
			&kainMaster.PanjangKain, &kainMaster.UpdateDate, &kainMaster.Status, &kainMaster.CreatedDate, &kainMaster.ModifiedDate,
			&kainMaster.CreatedBy, &kainMaster.ModifiedBy,
		)
		helper.PanicIfError(err)
		weavingKainMasters = append(weavingKainMasters, kainMaster)
	}
	return weavingKainMasters
}
