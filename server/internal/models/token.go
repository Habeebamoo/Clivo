package models

type TokenPayload struct {
	UserId  string  `json:"userId"`
	Role    string  `json:"role"`
}