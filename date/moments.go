package helpers

import "time"

// BeginningOfTheDay returns moment when this day begins - 00:00:00
func BeginningOfTheDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 1, time.Local)

}

// EndOfTheDay returns moment when this day ends - 23:59:59
func EndOfTheDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 1, time.Local)
}

// BeginningOfMonth returns moment when this month begins
func BeginningOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 0, -date.Day()+1)
}

// EndOfMonth returns moment when this month is ends
func EndOfMonth(date time.Time) time.Time {
	end := date.AddDate(0, 1, -date.Day())
	return time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, 0, time.Local)
}

// BeginningOfThisWeek returns beginning of this week
func BeginningOfThisWeek(when time.Time) time.Time {
	year, week := when.ISOWeek()
	date := time.Date(when.Year(), 0, 0, 0, 0, 0, 0, when.Location())
	isoYear, isoWeek := date.ISOWeek()
	for date.Weekday() != time.Monday { // iterate back to Monday
		date = date.AddDate(0, 0, -1)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoYear < year { // iterate forward to the first day of the first week
		date = date.AddDate(0, 0, 1)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoWeek < week { // iterate forward to the first day of the given week
		date = date.AddDate(0, 0, 1)
		isoYear, isoWeek = date.ISOWeek()
	}
	return date
}
