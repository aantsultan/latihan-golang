package helper

import (
	"latihan-golang-restful-api-nkw/dto"
	"latihan-golang-restful-api-nkw/model"
)

func ToFactoryMasterDto(factoryMaster model.FactoryMaster) dto.FactoryMasterDto {
	createdDate := ""
	if factoryMaster.CreatedDate.Valid {
		createdDate = factoryMaster.CreatedDate.Time.String()
	}
	modifiedDate := ""
	if factoryMaster.CreatedDate.Valid {
		modifiedDate = factoryMaster.ModifiedDate.Time.String()
	}
	var createdBy int64 = 0
	if factoryMaster.CreatedBy.Valid {
		createdBy = factoryMaster.CreatedBy.Int64
	}
	var modifiedBy int64 = 0
	if factoryMaster.ModifiedBy.Valid {
		createdBy = factoryMaster.ModifiedBy.Int64
	}

	return dto.FactoryMasterDto{
		Id:           factoryMaster.Id,
		Name:         factoryMaster.Name,
		Type:         factoryMaster.Type,
		Status:       factoryMaster.Status,
		CreatedDate:  createdDate,
		ModifiedDate: modifiedDate,
		CreatedBy:    createdBy,
		ModifiedBy:   modifiedBy,
	}
}

func ToFactoryMasterDtos(factoryMasters []model.FactoryMaster) []dto.FactoryMasterDto {
	var factoryMasterDtos []dto.FactoryMasterDto
	for _, factoryMaster := range factoryMasters {
		factoryMasterDtos = append(factoryMasterDtos, ToFactoryMasterDto(factoryMaster))
	}

	return factoryMasterDtos
}

func ToKainMasterDto(kainMaster model.WeavingKainMaster) dto.WeavingKainMasterDto {
	createdDate := ""
	if kainMaster.CreatedDate.Valid {
		createdDate = kainMaster.CreatedDate.Time.String()
	}
	modifiedDate := ""
	if kainMaster.CreatedDate.Valid {
		modifiedDate = kainMaster.ModifiedDate.Time.String()
	}
	var createdBy int64 = 0
	if kainMaster.CreatedBy.Valid {
		createdBy = kainMaster.CreatedBy.Int64
	}
	var modifiedBy int64 = 0
	if kainMaster.ModifiedBy.Valid {
		createdBy = kainMaster.ModifiedBy.Int64
	}
	updateDate := ""
	if kainMaster.UpdateDate.Valid {
		updateDate = kainMaster.UpdateDate.Time.String()
	}

	return dto.WeavingKainMasterDto{
		Id:           kainMaster.Id,
		KainName:     kainMaster.KainName,
		Anyaman:      kainMaster.Anyaman,
		ActiveStatus: kainMaster.ActiveStatus,
		LebarKain:    kainMaster.LebarKain,
		PanjangKain:  kainMaster.PanjangKain,
		UpdateDate:   updateDate,
		Status:       kainMaster.Status,
		CreatedDate:  createdDate,
		ModifiedDate: modifiedDate,
		CreatedBy:    createdBy,
		ModifiedBy:   modifiedBy,
	}
}

func ToWeavingKainMasterDtos(kainMasters []model.WeavingKainMaster) []dto.WeavingKainMasterDto {
	var kainMasterDtos []dto.WeavingKainMasterDto
	for _, kainMaster := range kainMasters {
		kainMasterDtos = append(kainMasterDtos, ToKainMasterDto(kainMaster))
	}

	return kainMasterDtos
}
