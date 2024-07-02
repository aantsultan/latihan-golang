package service

import (
	"context"
	"latihan-golang-restful-api-nkw/dto"
)

type FactoryMasterService interface {
	FindAll(ctx context.Context) dto.WebResponseDto[[]dto.FactoryMasterDto]
}
