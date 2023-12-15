package entities

import "time"

type Transaction struct {
	From_user_id int
	To_user_id   int
	Amount       int
	Status       string
	Message      string
	Created_at   time.Time
}
