package latihan_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	Name      string
	Email     string
	Age       int
	IsMarried bool
	Hobbies   []string
	Addresses []Address
}

type Address struct {
	Street string
	City   string
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		Name:      "Aant Sultan Rahmanya",
		Email:     "aant@mail.com",
		Age:       10,
		IsMarried: false,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
