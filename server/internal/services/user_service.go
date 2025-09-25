package services

import (
	"github.com/Habeebamoo/Clivo/server/internal/models"
	"github.com/Habeebamoo/Clivo/server/internal/repositories"
	"github.com/Habeebamoo/Clivo/server/pkg/utils"
)

type UserService interface {
	SignInUser(models.User) (string, int, error)
}

type UserSvc struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &UserSvc{repo}
}

func (uSvc *UserSvc) SignInUser(user models.User) (string, int, error) {
	//checks if user already exists and return jwt token
	exists := uSvc.repo.UserExists(user.Email) 

	if exists {
		//get user
		foundUser, code, err := uSvc.repo.GetUserByEmail(user.Email)
		if err != nil {
			return "", code, err
		}

		//sign jwt
		token, err := utils.SignToken(models.TokenPayload{ UserId: foundUser.UserId, Role: foundUser.Role })
		if err != nil {
			return "", 401, err
		}

		return token, 200, nil
	}

	//user was not found
	//now creating user

	createdUser, statusCode, err := uSvc.repo.CreateUser(user)
	if err != nil {
		return "", statusCode, err
	}

	//sign jwt token
	token, err := utils.SignToken(models.TokenPayload{ UserId: createdUser.UserId, Role: createdUser.Role })
	if err != nil {
		return "", 401, err
	}

	return token, 200, nil
}