package models

import (
	"time"
)

type Step struct {
	User     string
	Date     time.Time
	Week     int
	Year     int
	Goal     int
	Progress int
}
