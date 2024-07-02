package latihan_golang_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	create, _ := os.Create("resources/customer_out.json")
	encoder := json.NewEncoder(create)

	customer := Customer{
		Name:  "aant",
		Email: "aant@mail.com",
		Age:   20,
	}

	err := encoder.Encode(customer)
	if err != nil {
		return
	}
}
