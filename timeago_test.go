package timeago

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	tm := time.Now()
	s, err := TimeAgo(tm)
	if err != nil {
		t.Error(err)
	}
	if s != "right now" {
		t.Error("Expected: 'right now' but got ", s)
	}
}

func TestFewSeconds(t *testing.T) {
	ago := time.Now().Add(-time.Second * 4)
	s, err := TimeAgo(ago)
	if err != nil {
		t.Error(err)
	}
	if s != "a few seconds ago" {
		t.Error("Expected: 'a few seconds ago' but got ", s)
	}
}

func TestSeconds(t *testing.T) {
	ago := time.Now().Add(-time.Second * 15)
	s, err := TimeAgo(ago)
	if err != nil {
		t.Error(err)
	}
	if s != "15 seconds ago" {
		t.Error("Expected: '15 seconds ago' but got ", s)
	}
}
