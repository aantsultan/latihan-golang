package latihan_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {
	customer := Customer{
		Name:    "aant",
		Email:   "mail@mail.com",
		Hobbies: []string{"game", "coding", "manga"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{"Name":"aant","Email":"mail@mail.com","Age":0,"IsMarried":false,"Hobbies":["game","coding","manga"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.Name)
	fmt.Println(customer.Hobbies)
}

func TestJsonArrayComplex(t *testing.T) {
	customer := Customer{
		Name:    "aant",
		Email:   "mail@mail.com",
		Hobbies: []string{"game", "coding", "manga"},
		Addresses: []Address{{
			Street: "subam",
			City:   "jakut",
		}, {
			Street: "perum",
			City:   "bekasi",
		}},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonString := `{"Name":"aant","Email":"mail@mail.com","Age":0,"IsMarried":false,"Hobbies":["game","coding","manga"],"Addresses":[{"Street":"subam","City":"jakut"},{"Street":"perum","City":"bekasi"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.Name)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Addresses)
}

func TestJsonArrayComplexDecodeOnlyJsonArray(t *testing.T) {
	jsonString := `[{"Street":"subam","City":"jakut"},{"Street":"perum","City":"bekasi"}]`
	jsonBytes := []byte(jsonString)

	var addresses []Address

	err := json.Unmarshal(jsonBytes, &addresses)
	if err != nil {
		panic(err)
	}
	fmt.Println(addresses)
}

func TestJsonArrayComplexOnlyArray(t *testing.T) {
	address := []Address{{
		Street: "subam",
		City:   "jakut",
	}, {
		Street: "perum",
		City:   "bekasi",
	}}

	bytes, _ := json.Marshal(address)
	fmt.Println(string(bytes))
}
