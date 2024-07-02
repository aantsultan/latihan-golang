package latihan_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	jsonString := `{"id":"product-id", "name":"asus zenfone", "price":200000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "product-id",
		"name":  "product name",
		"price": 200000,
	}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

type Manager struct {
	Fullname    string `json:"full_name,omitempty"`
	Position    string `json:"position,omitempty"`
	Age         any    `json:"age"`
	YearsInWork any    `json:"years_in_work,omitempty"`
}

func TestMapEncodeLogic(t *testing.T) {
	manager := Manager{
		Fullname:    "FullnameFullnameFullname",
		Position:    "manager",
		Age:         10,
		YearsInWork: 101,
	}

	if !(len(manager.Fullname) >= 1) || !(len(manager.Fullname) <= 15) {
		manager.Fullname = ""
	}
	if !(len(manager.Position) >= 1) || !(len(manager.Position) <= 100) {
		manager.Position = ""
	}
	if !((int32)(manager.YearsInWork.(int)) >= 1) || !((int32)(manager.YearsInWork.(int)) <= 100) {
		manager.YearsInWork = nil
	}
	result, err := json.Marshal(manager)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
}
