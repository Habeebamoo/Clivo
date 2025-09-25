package services

import (
	"github.com/Habeebamoo/Clivo/server/internal/models"
	"github.com/Habeebamoo/Clivo/server/internal/repositories"
	"github.com/Habeebamoo/Clivo/server/pkg/utils"
)

type AuthService interface {
	SignInUser(models.User) (string, int, error)
}

type AuthSvc struct {
	repo repositories.AuthRepository
}

func NewAuthService(repo repositories.AuthRepository) AuthService {
	return &AuthSvc{repo}
}

func (as *AuthSvc) SignInUser(user models.User) (string, int, error) {
	//checks if user already exists and return jwt token
	exists := as.repo.UserExists(user.Email) 

	if exists {
		//get user
		foundUser, code, err := as.repo.GetUserByEmail(user.Email)
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

	createdUser, statusCode, err := as.repo.CreateUser(user)
	if err != nil {
		return "", statusCode, err
	}

	//sign jwt token
	token, err := utils.SignToken(models.TokenPayload{ UserId: createdUser.UserId, Role: createdUser.Role })
	if err != nil {
		return "", 401, err
	}

	return token, 200, nil

	//send email notification to admin
}