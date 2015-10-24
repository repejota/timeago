package timeago

import (
	"fmt"
	"time"
)

// TimeAgo ...
func TimeAgo(from time.Time) (s string) {

	duration := time.Now().Sub(from)
	fmt.Println(duration.Seconds())
	/*
		if years > 1 {
			fmt.Println("foooobar")
		}
		syears := strconv.FormatFloat(years, 'f', -1, 64)
		return string(syears)
	*/
	return ""
}
