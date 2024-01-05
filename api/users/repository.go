package users

import (
	"errors"

	"gorm.io/gorm"
)

var ErrDatabaseConnectionNil = errors.New("Database connection is nil")

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) FindAll() ([]User, error) {
	var users []User
	if ur.db == nil {
		return users, ErrDatabaseConnectionNil
	}
	ur.db.Find(&users)
	return users, nil
}

func (ur *UserRepository) FindByID(id string) (User, error) {
	var user User
	if ur.db == nil {
		return user, ErrDatabaseConnectionNil
	}
	if err := ur.db.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (ur *UserRepository) Create(user *User) error {
	if ur.db == nil {
		return ErrDatabaseConnectionNil
	}
	ur.db.Create(user)
	return nil
}

func (ur *UserRepository) UpdateByID(id string, updatedUser *User) error {
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

func (ur *UserRepository) DeleteByID(id string) error {
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
