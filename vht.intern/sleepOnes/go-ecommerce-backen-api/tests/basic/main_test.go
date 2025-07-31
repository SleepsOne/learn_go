package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {
	var (
		input = 1
		want  = 2
	)
	actual := AddOne(input)
	if actual != want {
		t.Errorf("AddOne(%d) = %d; want %d", input, actual, want)
	}

}

func TestAssert(t *testing.T) {
	assert.Equal(t, 1, 2, "they should be equal")
	fmt.Println("TestAssert passed")
}

func TestRequire(t *testing.T) {
	require.Equal(t, 1, 2, "they should be equal")
	fmt.Println("TestRequire passed")
}
