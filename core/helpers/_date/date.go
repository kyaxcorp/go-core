package _date

import (
	"time"
)

// RangeDate returns a date range function over start date to end date inclusive.
// After the end of the range, the range function returns a zero date,
// date.IsZero() is true.
func RangeDate(start, end time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date
	}
}

type DailyFormattedInput struct {
	Start time.Time
	End   time.Time
	Cb    DailyFormattedCallback
}

type DailyFormattedCallback func(dt string) string

func RangeDateDailyFormatted(input DailyFormattedInput) []string {
	var formatted []string
	for rd := RangeDate(input.Start, input.End); ; {
		date := rd()
		if date.IsZero() {
			break
		}
		dt := date.Format("2006-01-02")
		if input.Cb != nil {
			// Use the callback
			dt = input.Cb(dt)
		}
		formatted = append(formatted, dt)
	}
	return formatted
}
