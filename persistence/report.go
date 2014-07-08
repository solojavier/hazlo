package persistence

import (
	"time"
)

type Report struct {
	User        string
	Date        time.Time
	Week        int
	Year        int
	Goal        int
	Progress    int
	Measurement string
}
