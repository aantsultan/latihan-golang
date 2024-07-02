package service

import (
	"context"
	"database/sql"
	"latihan-golang-restful-api-nkw/dto"
	"latihan-golang-restful-api-nkw/helper"
	"latihan-golang-restful-api-nkw/repository"
	"net/http"
)

type WeavingKainMasterServiceImpl struct {
	KainMasterRepository repository.WeavingKainMasterRepository
	DB                   *sql.DB
}

func NewWeavingKainMasterService(kainMasterRepository repository.WeavingKainMasterRepository, DB *sql.DB) *WeavingKainMasterServiceImpl {
	return &WeavingKainMasterServiceImpl{
		KainMasterRepository: kainMasterRepository,
		DB:                   DB,
	}
}

func (service WeavingKainMasterServiceImpl) FindAll(ctx context.Context) dto.WebResponseDto[[]dto.WeavingKainMasterDto] {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	kainMasters := service.KainMasterRepository.FindAll(ctx, tx)
	kainMasterDtos := helper.ToWeavingKainMasterDtos(kainMasters)
	return dto.WebResponseDto[[]dto.WeavingKainMasterDto]{
		Data:    kainMasterDtos,
		Status:  http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}
}
