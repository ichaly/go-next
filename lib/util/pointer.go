package util

import (
	"time"
)

/**
 * https://github.com/pilagod/pointer
 */

func TimePtr(t time.Time) *time.Time {
	return &t
}

func StringPtr(s string) *string {
	return &s
}