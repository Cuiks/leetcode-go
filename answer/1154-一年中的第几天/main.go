package main

import (
	"time"
)

func dayOfYear(date string) int {
	if date == "" {
		return 0
	}
	t, _ := time.Parse("2006-01-02", date)
	return t.YearDay()
}

func main() {
	dayOfYear("")
}
