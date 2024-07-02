package dto

type WeavingKainMasterDto struct {
	Id           int64   `json:"id"`
	KainName     string  `json:"kainName"`
	Anyaman      string  `json:"anyaman"`
	ActiveStatus bool    `json:"activeStatus"`
	LebarKain    float64 `json:"lebarKain"`
	PanjangKain  float64 `json:"panjangKain"`
	UpdateDate   string  `json:"updateDate"`
	Status       bool    `json:"status"`
	CreatedDate  string  `json:"createdDate"`
	ModifiedDate string  `json:"modifiedDate"`
	CreatedBy    int64   `json:"createdBy"`
	ModifiedBy   int64   `json:"modifiedBy"`
}
