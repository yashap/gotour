package model

// MyFloat is just a float64 with some extra methods
type MyFloat float64

// Abs returns the absolute value of the float
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
