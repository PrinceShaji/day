package day

import (
	"errors"
	"time"
)

type DayInterface interface{}

type Day struct{}

func GetLayout(layout string) string {

	return generateLayout(layout)
}

// String-parse-time
// Parse time string using layout
func Strptime(layout string, value string) (time.Time, error) {
	layout = generateLayout(layout)

	return time.Parse(layout, value)
}

// String-from-time
// Generate datetime string from time.Time or day.Day using layout
func Strftime(layout string, value DayInterface) (resp string, err error) {
	layout = generateLayout(layout)

	switch v := value.(type) {
	case time.Time:
		return v.Format(layout), nil
	case Day:
		return "", nil
	default:
		err = errors.New("wrong type")
		return
	}
}
