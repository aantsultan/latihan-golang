package test

import (
	"github.com/stretchr/testify/assert"
	"latihan-golang-restful-api-nkw/helper"
	"testing"
)

func TestProperties(t *testing.T) {
	value := helper.GetProperties("hello.world")
	assert.Equal(t, "hello world", value)

	username := helper.GetProperties("datasource.username")
	password := helper.GetProperties("datasource.password")
	assert.Equal(t, "nikawaweb", username)
	assert.Equal(t, "Sabeso76", password)
}
