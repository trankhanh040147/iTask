package utils

import (
	"fmt"
	"time"
)

func ParseStringToTime(date string) (*time.Time, error) {
	layout := "02-01-2006"
	dateRes, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return nil, nil
	}
	return &dateRes, nil
}

func ParseTimeToString(date *time.Time) string {
	layout := "02-01-2006"

	formattedTime := date.Format(layout)
	return formattedTime
}
