package hello_world_test

import (
	helloworld "go-tdd"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	name := "Nathan"
	greeting := helloworld.SayHello(name)

	assert.Equal(t, "Hello Nathan", greeting)
}
