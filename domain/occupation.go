package domain

import "time"

type Occupation struct {
	ID       uint
	Name     string
	Position string
	Start    time.Time
	End      time.Time
	URL      string
}
