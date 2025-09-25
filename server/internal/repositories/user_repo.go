package repositories

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Habeebamoo/Clivo/server/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(models.User) (models.User, int, error)
	GetUserByEmail(email string) (models.User, int, error)
	UserExists(email string) bool
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepo{db: db}
}

func (uRepo *UserRepo) CreateUser(user models.User) (models.User, int, error) {
	res := uRepo.db.Create(&user)

	if res.Error != nil {
		if strings.Contains(res.Error.Error(), "duplicate key value") {
			return models.User{}, http.StatusConflict, fmt.Errorf("user already exist")
		}
		return models.User{}, 500, fmt.Errorf("internal server error")
	}

	return user, 201, nil
}

func (uRepo *UserRepo) GetUserByEmail(email string) (models.User, int, error) {
	var user models.User
	res := uRepo.db.First(&user, "email = ?", email)
	
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return user, 404, fmt.Errorf("user not found")
		}
		return user, 500, fmt.Errorf("internal server error")
	}

	return user, 200, nil
}

func (uRepo *UserRepo) UserExists(email string) bool {
	var user models.User
	res := uRepo.db.First(&user, "email = ?", email)
	
	//no user
	if res.Error == gorm.ErrRecordNotFound {
		return false
	}

	return true
}