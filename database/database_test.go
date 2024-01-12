package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetName(t *testing.T) {
	name := GetName("John")
	assert.Equal(t, "John", name)
}

func TestGetAge(t *testing.T) {
	age := GetAge(20)
	assert.Equal(t, 20, age)
}

func TestGetHelloName(t *testing.T) {
	hello := GetHello("John")
	assert.Equal(t, "Hello John", hello)
}
