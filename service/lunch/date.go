package lunch

import "time"

func truncateDate(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

func startOfWeek(date time.Time) time.Time {
	return truncateDate(date.AddDate(0, 0, -int(date.Weekday())))
}

func endOfWeek(date time.Time) time.Time {
	return startOfWeek(date).AddDate(0, 0, 7)
}
