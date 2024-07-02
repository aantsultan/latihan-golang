package latihan_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestJsonTag(t *testing.T) {
	product := Product{
		Id:       "productId",
		Name:     "product name",
		ImageUrl: "/image/url",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJsonTagDecode(t *testing.T) {
	jsonString := `{"id":"productId","name":"product name","image_url":"/image/url"}`
	jsonBytes := []byte(jsonString)

	product := Product{}
	json.Unmarshal(jsonBytes, &product)
	fmt.Println(product)
	fmt.Println(product.Id)
	fmt.Println(product.Name)
	fmt.Println(product.ImageUrl)
}
