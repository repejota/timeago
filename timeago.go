package timeago

import "time"

// Seconds is a constant representing the number of seconds as a base
// calculation unit
const Second = 1

// Minute is the number of seconds in a minute
const Minute = Second * 60

// Hour is the number of seconds in a hour
const Hour = Minute * 60

// Day is the number of seconds in a day
const Day = Hour * 24

// Week is the number of seconds in a week
const Week = Day * 7

// Month is the number of seconds in a month
const Month = Day * 30

// Year is the number of seconds in a year
const Year = Day * 365

// TimeAgo ...
func TimeAgo(from time.Time) (s string) {
	/*
		duration := time.Now().Sub(from)
		years := duration.Hours() / 24 / 365

			if years > 1 {
				fmt.Println("foooobar")
			}
			syears := strconv.FormatFloat(years, 'f', -1, 64)
			return string(syears)
	*/
	return ""
}
