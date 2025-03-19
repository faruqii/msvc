package repository

import (
	"github.com/faruqii/msvc/authservice/internal/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(User *entities.User) (err error)
	GetByEmail(email string) (user *entities.User, err error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepositoryImpl {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) Create(user *entities.User) (err error) {
	return r.db.Create(user).Error
}

func (r *userRepositoryImpl) GetByEmail(email string) (user *entities.User, err error) {
	user = &entities.User{}
	err = r.db.Where("email = ?", email).First(user).Error
	return
}
