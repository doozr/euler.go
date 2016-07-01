package benchmark

import "time"

func Start() time.Time {
	return time.Now()
}

func End(start time.Time) int {
	duration := time.Since(start)
	return int(duration.Nanoseconds() / 1000)
}
