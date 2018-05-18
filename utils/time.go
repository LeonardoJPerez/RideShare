package utils

import (
	"strconv"
	"time"
)

// BeginningOfDay returns the Time object representing the start of the day
func BeginningOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// EndOfDay returns the Time object representing the end of the day
func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, t.Location())
}

// BeginningOfToday returns the Time object representing the start of the current day
func BeginningOfToday() time.Time {
	return BeginningOfDay(time.Now())
}

// EndOfToday returns the Time object representing the end of the current day
func EndOfToday() time.Time {
	return EndOfDay(time.Now())
}

// ParseEpoch converts an epoch string into a time.Time object.
func ParseEpoch(epoch string) (*time.Time, error) {
	i, err := strconv.ParseInt(epoch, 10, 64)
	if err != nil {
		return nil, err
	}

	tm := time.Unix(i, 0)
	return &tm, nil
}
