package models

import (
	"time"
)

type Activity struct {
	ID     int64
	Symbol string
	Start  time.Time
	Note   string
}

type Temperature struct {
	Min int
	Now int
	Max int
}
