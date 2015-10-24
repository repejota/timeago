package timeago

import (
	"errors"
	"strconv"
	"time"
)

// TimeAgo ...
func TimeAgo(from time.Time) (s string, err error) {

	duration := time.Now().Sub(from)

	// right now
	if duration.Seconds() < 1 {
		return "right now", nil
	}

	// a few seconds ago
	if duration.Seconds() < 15 {
		return "a few seconds ago", nil
	}

	// x seconds ago
	if duration.Seconds() < 60 {
		return strconv.Itoa(int(duration.Seconds())) + " seconds ago", nil
	}

	// unknown error
	err = errors.New("Unknown error")
	return "", err
}
