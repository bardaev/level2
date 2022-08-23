package main

import "time"

func GetDate(date string) (time.Time, error) {
	return time.Parse("2-1-2006", date)
}
