package test

import (
	"github.com/stretchr/testify/assert"
	"latihan-golang-depedency-injection/simple"
	"testing"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("database")
	assert.NotNil(t, connection)
	cleanup()
}
