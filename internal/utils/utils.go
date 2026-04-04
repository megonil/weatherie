// Package utils
package utils

import "time"

func Seconds(seconds uint) time.Duration {
	return time.Second * time.Duration(seconds)
}
