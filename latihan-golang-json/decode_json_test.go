package latihan_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"Name":"Aant Sultan Rahmanya","Email":"aant@mail.com","Age":10,"IsMarried":false}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.Name)
	fmt.Println(customer.Email)
	fmt.Println(customer.Age)
	fmt.Println(customer.IsMarried)

}
