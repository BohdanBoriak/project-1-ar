package domain

import "time"

type User struct {
	Id   uint64        `json:"id"`
	Name string        `json:"name"`
	Time time.Duration `json:"time"`
}
