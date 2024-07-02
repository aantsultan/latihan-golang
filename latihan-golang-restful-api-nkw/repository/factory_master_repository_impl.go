package repository

import (
	"context"
	"database/sql"
	"latihan-golang-restful-api-nkw/helper"
	"latihan-golang-restful-api-nkw/model"
)

type FactoryMasterRepositoryImpl struct {
}

func NewFactoryMasterRepository() *FactoryMasterRepositoryImpl {
	return &FactoryMasterRepositoryImpl{}
}

func (repository FactoryMasterRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []model.FactoryMaster {
	//TODO implement me
	SQL := "select factory_id, factory_name, factory_type, status, created_date, modified_date, created_by, modified_by from s_factory_master "
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var factoryMasters []model.FactoryMaster
	for rows.Next() {
		factoryMaster := model.FactoryMaster{}
		err := rows.Scan(&factoryMaster.Id, &factoryMaster.Name, &factoryMaster.Type, &factoryMaster.Status,
			&factoryMaster.CreatedDate, &factoryMaster.ModifiedDate, &factoryMaster.CreatedBy, &factoryMaster.ModifiedBy)
		helper.PanicIfError(err)
		factoryMasters = append(factoryMasters, factoryMaster)
	}
	return factoryMasters
}
