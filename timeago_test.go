package timeago

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	tm := time.Now()
	s := TimeAgo(tm)
	if s != "right now" {
		t.Error("Expected: 'right now' but got ", s)
	}
}

func TestFewSeconds(t *testing.T) {
	ago := time.Unix(time.Now().Unix()-4, 0)
	s := TimeAgo(ago)
	if s != "few seconds ago" {
		t.Error("Expected: 'few seconds ago' but got ", s)
	}
}
