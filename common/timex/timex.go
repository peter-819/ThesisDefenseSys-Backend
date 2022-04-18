package timex

import (
	"time"
)

func TimeToString(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.String()
}

func GStringToTime(t string) (time.Time, error) {
	if t == "" {
		return time.Time{}, nil
	}
	return time.Parse("2006-01-02 15:04:05 -0700 MST", t)
}

func FStringToTime(t string) (time.Time, error) {
	if t == "" {
		return time.Time{}, nil
	}
	return time.Parse("2006-01-02T15:04:05", t)
}
