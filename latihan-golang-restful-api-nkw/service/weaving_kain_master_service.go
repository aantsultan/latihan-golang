package service

import (
	"context"
	"latihan-golang-restful-api-nkw/dto"
)

type WeavingKainMasterService interface {
	FindAll(ctx context.Context) dto.WebResponseDto[[]dto.WeavingKainMasterDto]
}
