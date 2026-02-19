package domain

import "time"

type Event struct {
	ID   int
	Name string
	Date time.Time
}
