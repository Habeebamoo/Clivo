package models

import "time"

type User struct {
	UserId     string     `json:"userId"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Username   string     `json:"username"`
	Role       string     `json:"role"`
	CreatedAt  time.Time  `json:"createdAt"`
}

type Profile struct {
	UserId     string  `json:"userId"`
	Picture    string  `json:"picture"`
	Following  int     `json:"following"`
	Followers  int     `json:"followers"`
}