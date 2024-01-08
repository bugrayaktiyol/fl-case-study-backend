package users

import (
	"errors"

	"gorm.io/gorm"
)

//go:generate mockery --name UserRepository
type UserRepository interface {
	FindAll() ([]*User, error)
	FindByID(id string) (User, error)
	Create(user *User) error
	UpdateByID(id string, updatedUser *User) error
	DeleteByID(id string) error
}

var ErrDatabaseConnectionNil = errors.New("Database connection is nil")

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) FindAll() ([]*User, error) {
	var users []*User
	if ur.db == nil {
		return users, ErrDatabaseConnectionNil
	}
	ur.db.Find(&users)
	return users, nil
}

func (ur *userRepository) FindByID(id string) (User, error) {
	var user User
	if ur.db == nil {
		return user, ErrDatabaseConnectionNil
	}
	if err := ur.db.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (ur *userRepository) Create(user *User) error {
	if ur.db == nil {
		return ErrDatabaseConnectionNil
	}
	ur.db.Create(user)
	return nil
}

func (ur *userRepository) UpdateByID(id string, updatedUser *User) error {
	var user User
	if ur.db == nil {
		return ErrDatabaseConnectionNil
	}
	if err := ur.db.First(&user, id).Error; err != nil {
		return err
	}
	ur.db.Model(&user).Updates(updatedUser)
	return nil
}

func (ur *userRepository) DeleteByID(id string) error {
	var user User
	if ur.db == nil {
		return ErrDatabaseConnectionNil
	}
	if err := ur.db.First(&user, id).Error; err != nil {
		return err
	}
	ur.db.Delete(&user)
	return nil
}
