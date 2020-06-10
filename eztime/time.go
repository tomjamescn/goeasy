package eztime

import (
	"fmt"
	"time"
)

var TimeFormat = "2006-01-02"

func MustTime(timeStr string) time.Time {
	if t, err := time.Parse(TimeFormat, timeStr); err != nil {
		panic(err)
	} else {
		return t
	}
}

func MustDate(year, month, day int) time.Time {
	if t, err := time.Parse(TimeFormat, fmt.Sprintf("%d-%02d-%02d", year, month, day)); err != nil {
		panic(err)
	} else {
		return t
	}
}

func FormatDate(t time.Time) string {
	return t.Format(TimeFormat)
}
