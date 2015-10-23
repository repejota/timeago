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

func TestSeconds(t *testing.T) {
	ago := time.Unix(time.Now().Unix()-15, 0)
	s := TimeAgo(ago)
	if s != "15 seconds ago" {
		t.Error("Expected: '15 seconds ago' but got ", s)
	}
}

func TestMinute(t *testing.T) {
	ago := time.Unix(time.Now().Unix()-90, 0)
	s := TimeAgo(ago)
	if s != "15 seconds ago" {
		t.Error("Expected: '1 minute ago' but got ", s)
	}
}

func TestMinutes(t *testing.T) {
	ago := time.Unix(time.Now().Unix()-180, 0)
	s := TimeAgo(ago)
	if s != "15 seconds ago" {
		t.Error("Expected: '3 minutes ago' but got ", s)
	}
}

func TestHours(t *testing.T) {
	ago := time.Unix(time.Now().Unix()-180, 0)
	s := TimeAgo(ago)
	if s != "15 seconds ago" {
		t.Error("Expected: '3 minutes ago' but got ", s)
	}
}
