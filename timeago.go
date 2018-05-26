package timeago

import (
	"errors"
	"strconv"
	"time"
)

// TimeAgo ...
type TimeAgo struct {
	duration time.Duration
}

// NewTimeAgo ...
func NewTimeAgo(from time.Time) *TimeAgo {
	timeago := &TimeAgo{}
	timeago.duration = time.Now().Sub(from)
	return timeago
}

// Render ...
func (t *TimeAgo) Render() (string, error) {
	// right now
	if t.duration.Seconds() < 1 {
		return "right now", nil
	}

	// a few seconds ago
	if t.duration.Seconds() < 15 {
		return "a few seconds ago", nil
	}

	// x seconds ago
	if t.duration.Seconds() < 60 {
		return strconv.Itoa(int(t.duration.Seconds())) + " seconds ago", nil
	}

	// unknown error
	err := errors.New("Unknown error")

	return err.Error(), err
}
