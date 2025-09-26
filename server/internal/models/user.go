package models

import "time"

type User struct {
	UserId     string     `json:"userId"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Role       string     `json:"role"`
	Verified   bool       `json:"verified"`
	CreatedAt  time.Time  `json:"createdAt"`
}

type Profile struct {
	UserId       string  `json:"userId"`
	Username     string  `json:"username"`
	Bio          string  `json:"bio"`
	Picture      string  `json:"picture"`
	ProfileLink  string  `json:"profileLink"`
	Following    int     `json:"following"`
	Followers    int     `json:"followers"`
}