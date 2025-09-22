package handlers

import "github.com/Habeebamoo/Clivo/server/internal/services"

type UserHandler struct {
	svc services.UserService
}

func NewUserHandler(svc services.UserService) UserHandler {
	return UserHandler{svc}
}
