package timeutils

import "fmt"

func SecondToClock(s int) string {
	seconds := s % 60
	minutes := (s - seconds) / 60

	secondsStr := fmt.Sprint(seconds)
	if len(secondsStr) == 1 {
		secondsStr = fmt.Sprint("0", seconds)
	}

	return fmt.Sprint(minutes, ":", secondsStr)
}
