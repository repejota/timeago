package timeago

import (
	"errors"
	"time"
)

// TimeAgo ...
func TimeAgo(from time.Time) (s string, err error) {

	duration := time.Now().Sub(from)
	if duration.Seconds() < 1 {
		return "right now", nil
	}
	err = errors.New("Unknown error")
	return "", err
}
