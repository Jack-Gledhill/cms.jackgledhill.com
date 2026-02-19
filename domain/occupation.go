package domain

import "time"

type Occupation struct {
	ID       int
	Name     string
	Position string
	Start    time.Time
	End      time.Time
	URL      string
}
