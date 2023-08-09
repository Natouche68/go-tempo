package timeutils

import "fmt"

func SecondToClock(s int) string {
	seconds := s % 60
	minutes := (s - seconds) / 60

	secondsStr := fmt.Sprint(seconds)
	if secondsStr == "0" {
		secondsStr = "00"
	}

	return fmt.Sprint(minutes, ":", secondsStr)
}
