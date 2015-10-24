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

func TestDay(t *testing.T) {
	ago := time.Now().AddDate(0, 0, -1).Add(-time.Second * 10)
	s := TimeAgo(ago)
	if s != "a day ago" {
		t.Error("Expected: 'a day ago' but got ", s)
	}
}

func TestDays(t *testing.T) {
	ago := time.Now().AddDate(0, 0, -4).Add(-time.Second * 10)
	s := TimeAgo(ago)
	if s != "4 days ago" {
		t.Error("Expected: '4 days ago' but got ", s)
	}
}

func TestWeek(t *testing.T) {
	ago := time.Now().AddDate(0, 0, -7).Add(-time.Second * 10)
	s := TimeAgo(ago)
	if s != "a week ago" {
		t.Error("Expected: 'a week ago' but got ", s)
	}
}

func TestWeeks(t *testing.T) {
	ago := time.Now().AddDate(0, 0, -16).Add(-time.Second * 10)
	s := TimeAgo(ago)
	if s != "2 weeks ago" {
		t.Error("Expected: '2 weeks ago' but got ", s)
	}
}

func TestMonth(t *testing.T) {
	ago := time.Now().AddDate(0, -1, 0).Add(-time.Second * 10)
	s := TimeAgo(ago)
	if s != "a month ago" {
		t.Error("Expected: 'a month ago' but got ", s)
	}
}

func TestMonths(t *testing.T) {
	ago := time.Now().AddDate(0, -5, 0).Add(-time.Second * 10)
	s := TimeAgo(ago)
	if s != "5 months ago" {
		t.Error("Expected: '5 months ago' but got ", s)
	}
}

func TestYear(t *testing.T) {
	ago := time.Now().AddDate(-1, 0, 0).Add(-time.Second * 10)
	s := TimeAgo(ago)
	if s != "a year ago" {
		t.Error("Expected: 'a year ago' but got ", s)
	}
}

func TestYears(t *testing.T) {
	ago := time.Now().AddDate(-2, 0, 0).Add(-time.Second * 10)
	s := TimeAgo(ago)
	if s != "2 years ago" {
		t.Error("Expected: '2 years ago' but got ", s)
	}
}
