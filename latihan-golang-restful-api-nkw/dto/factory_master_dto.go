package dto

type FactoryMasterDto struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Status       bool   `json:"status"`
	CreatedDate  string `json:"createdDate"`
	ModifiedDate string `json:"modifiedDate"`
	CreatedBy    int64  `json:"createdBy"`
	ModifiedBy   int64  `json:"modifiedBy"`
}
