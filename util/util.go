package util

import "time"

func TimeNow() *time.Time {
	t := time.Now()
	return &t
}

