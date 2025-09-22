package repositories

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser()
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepo{db: db}
}

func (uRepo *UserRepo) CreateUser() {

}