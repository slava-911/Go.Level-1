package circle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDiameterBySquare(t *testing.T) {
	d, err := DiameterBySquare(10)
	require.NoError(t, err)
	assert.EqualValues(t, 3.568248, d)
}

func TestLengthBySquare(t *testing.T) {
	l, err := LengthBySquare(10)
	require.NoError(t, err)
	require.Equal(t, 11.209982, l)
}

func ExampleDiameterBySquare() {
	d, _ := DiameterBySquare(10)
	fmt.Printf("%f\n", d)

	// Output:
	// 3.568248
}

func ExampleLengthBySquare() {
	l, _ := LengthBySquare(10)
	fmt.Printf("%f\n", l)

	// Output:
	// 11.209982
}
