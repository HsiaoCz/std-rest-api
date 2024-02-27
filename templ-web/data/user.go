package data

import "time"

type User struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Vatar     string `json:"vartar"`
	Identity  string `json:"identity"`
	CreatedAt time.Time
}
