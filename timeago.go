package main

import "time"

const (
	// Day is the duratino of a day
	Day time.Duration = time.Hour * 24
	// Month is the duration of a month
	Month time.Duration = Day * 30
	// Year is the duration of a year
	Year time.Duration = Day * 365
)
