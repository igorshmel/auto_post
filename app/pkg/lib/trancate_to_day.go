package lib

import "time"

// TruncateToDay --
func TruncateToDay(date time.Time) int64 {
	year, month, day := date.Date()
	truncatedDate := time.Date(year, month, day, 0, 0, 0, 0, date.Location())
	return truncatedDate.Unix()
}
