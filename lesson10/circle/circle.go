package circle

import (
	"errors"
	"math"
)

var ErrNotPositive = errors.New("input number must be positive")

// DiameterBySquare calculates the diameter of a circle from its square
func DiameterBySquare(s float64) (float64, error) {
	if s < 0 {
		return 0, ErrNotPositive
	}
	d := 2 * math.Sqrt(s/math.Pi)
	return math.Round(d*1000000) / 1000000, nil
}

// LengthBySquare calculates the length of a circle from its square
func LengthBySquare(s float64) (float64, error) {
	d, err := DiameterBySquare(s)
	if err != nil {
		return 0, err
	}
	l := math.Pi * d
	return math.Round(l*1000000) / 1000000, nil
}
