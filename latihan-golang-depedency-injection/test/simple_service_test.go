package test

import (
	"github.com/stretchr/testify/assert"
	"latihan-golang-depedency-injection/simple"
	"testing"
)

func TestSimpleServiceError(t *testing.T) {
	result, err := simple.InitializedService(true)
	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func TestSimpleServiceSuccess(t *testing.T) {
	result, err := simple.InitializedService(false)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
