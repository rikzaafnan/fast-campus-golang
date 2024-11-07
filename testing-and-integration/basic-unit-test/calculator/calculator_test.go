package calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	expected := 3

	// if result != expected {
	// 	t.Errorf("Add(1, 2) should return 3 but got: %d", result)
	// }

	assert.Equal(t, expected, result, "Add(1,2) should return 3")
}

func TestDivideBy(t *testing.T) {
	result, err := DivideBy(10, 2)
	assert.NoError(t, err, "it should return error")

	assert.Equal(t, 5, result, "Divide by (10, 5) should return 5")

}

func TestDivideBy_BiggerNumber(t *testing.T) {
	result, err := DivideBy(2, 10)
	assert.Error(t, err, "it should return error")

	assert.Equal(t, 0, result, "Divide by (10, 5) should return 0 due to insufficient number")

}

func TestDivideBy_ZeroNumber(t *testing.T) {
	assert.NotPanics(t, func() {
		_, _ = DivideBy(0, 10)
	}, "it should not panic event if we sent a 0 number")
}

// bisa juga mengecek :
// equal
// not equal
// greater then
// less then
// range
//  not panic
