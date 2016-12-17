package triangle

import (
	"math"
)

const testVersion = 3

// Notice it returns this type.  Pick something suitable.
type Kind string

// Pick values for the following identifiers used by the test program.
const NaT = "NaT" // not a triangle
const Equ = "Equ" // equilateral
const Iso = "Iso" // isosceles
const Sca = "Sca" // scalene

// Determine kind from side
func KindFromSides(a, b, c float64) Kind {

	// Sanity
	// if a <= 0 || b <= 0 || c <= 0 {
	// 	return NaT
	// }
	switch {
	case math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1):
		// Guard against infinity
		return NaT
	case math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c):
		// Guard against NaN
		return NaT
	case a <= 0 || b <= 0 || c <= 0:
		// Guard against non positive values
		return NaT
	case b+c < a || a+c < b || a+b < c:
		// Guard against
		return NaT
	case a == b && b == c && a == c:
		// Properties of Equalateral Triangle
		return Equ
	case b == a || a == c || b == c:
		// Properties of Isosceles Triangle
		return Iso
	case b+c >= a && a+c >= b && a+b >= c:
		// Properties of Scalar Triangle
		return Sca
	default:
		// Default not a Triangle'
		return NaT
	}
}
