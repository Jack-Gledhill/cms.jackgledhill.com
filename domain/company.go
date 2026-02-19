package domain

import "time"

type Company struct {
	ID       int
	Name     string
	Position string
	Start    time.Time
	End      time.Time
}
