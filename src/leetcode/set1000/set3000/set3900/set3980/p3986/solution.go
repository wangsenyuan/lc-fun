package p3986

import "time"

func secondsBetweenTimes(startTime string, endTime string) int {
	t1, _ := time.Parse("15:04:05", startTime)
	t2, _ := time.Parse("15:04:05", endTime)
	return int(t2.Sub(t1).Seconds())
}
