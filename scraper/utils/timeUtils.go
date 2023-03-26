package utils

import (
	"time"
)

func GetDateTimeFromEpochMillis(timeInMillis int64) time.Time {
	return time.Unix(timeInMillis, 0)
}
