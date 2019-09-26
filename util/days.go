package util

import (
	"strconv"
	"time"
)

func Days(date time.Time) string {
	now := time.Now()
	d := date.Sub(now)
	days := strconv.FormatInt(int64(d.Hours()/24), 10)

	return days
}
