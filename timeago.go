package timeago

import (
	"fmt"
	"strconv"
	"time"
)

// MakeString returns the number of miliseconds since a given time
func MakeString(from time.Time) (s string) {
	duration := time.Now().Sub(from)
	years := duration.Hours() / 24 / 365
	if years > 1 {
		fmt.Println("foooobar")
	}
	syears := strconv.FormatFloat(years, 'f', -1, 64)
	return string(syears)
}

//
