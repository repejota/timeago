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
	ago := time.Now().Add(-time.Second * 4)
	s := TimeAgo(ago)
	if s != "few seconds ago" {
		t.Error("Expected: 'few seconds ago' but got ", s)
	}
}

func TestSeconds(t *testing.T) {
	ago := time.Now().Add(-time.Second * 15)
	s := TimeAgo(ago)
	if s != "15 seconds ago" {
		t.Error("Expected: '15 seconds ago' but got ", s)
	}
}

func TestMinute(t *testing.T) {
	ago := time.Now().Add(-time.Minute).Add(-time.Second * 10)
	s := TimeAgo(ago)
	if s != "a minute ago" {
		t.Error("Expected: 'a minute ago' but got ", s)
	}
}

func TestMinutes(t *testing.T) {
	ago := time.Now().Add(-time.Minute * 3).Add(-time.Second * 10)
	s := TimeAgo(ago)
	if s != "3 minutes ago" {
		t.Error("Expected: '3 minutes ago' but got ", s)
	}
}

func TestHour(t *testing.T) {
	ago := time.Now().Add(-time.Hour).Add(-time.Second * 10)
	s := TimeAgo(ago)
	if s != "an hour ago" {
		t.Error("Expected: 'an hour ago' but got ", s)
	}
}

func TestHours(t *testing.T) {
	ago := time.Now().Add(-time.Hour * 3).Add(-time.Second * 10)
	s := TimeAgo(ago)
	if s != "15 hours ago" {
		t.Error("Expected: '3 hours ago' but got ", s)
	}
}
