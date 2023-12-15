package entities

import "time"

type Top_Up struct {
	User_id    int
	Amount     uint64
	Status     string
	Created_at time.Time
}
