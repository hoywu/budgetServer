package utils

import "time"

func IsTimeValid(t string) bool {
	_, err := time.Parse(time.DateTime, t)
	if err != nil {
		return false
	}
	return true
}
