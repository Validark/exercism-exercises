// Clock stub file

// To use the right term, this is the package *clause*.
// You can document general stuff about the package here if you like.
package clock

import "fmt"

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

// Clock API as stub definitions.  No, it doesn't compile yet.
// More details and hints are in clock_test.go.

type Clock int // Complete the type definition.  Pick a suitable data type.

func New(hour, minute int) Clock {
	c := Clock((hour*60 + minute) % 1440)
	if c < 0 {
		c += 1440
	}
	return c
}

func (c Clock) String() string {
	hours := c / 60
	minutes := c % 60
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+minutes)
}

// Remember to delete all of the stub comments.
// They are just noise, and reviewers will complain.