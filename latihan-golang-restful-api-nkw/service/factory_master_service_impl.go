package service

import (
	"context"
	"database/sql"
	"latihan-golang-restful-api-nkw/dto"
	"latihan-golang-restful-api-nkw/helper"
	"latihan-golang-restful-api-nkw/repository"
	"net/http"
)

type FactoryMasterServiceImpl struct {
	FactoryMasterRepository repository.FactoryMasterRepository
	DB                      *sql.DB
}

func NewFactoryMasterService(factoryMasterRepository repository.FactoryMasterRepository, DB *sql.DB) *FactoryMasterServiceImpl {
	return &FactoryMasterServiceImpl{
		FactoryMasterRepository: factoryMasterRepository,
		DB:                      DB,
	}
}

func (service FactoryMasterServiceImpl) FindAll(ctx context.Context) dto.WebResponseDto[[]dto.FactoryMasterDto] {
	//TODO implement me
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	factoryMasters := service.FactoryMasterRepository.FindAll(ctx, tx)
	factoryMastersDto := helper.ToFactoryMasterDtos(factoryMasters)
	return dto.WebResponseDto[[]dto.FactoryMasterDto]{
		Data:    factoryMastersDto,
		Status:  http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}
}
