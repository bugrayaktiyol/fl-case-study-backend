package users_test

import (
	"errors"
	"testing"

	"github.com/bugrayaktiyol/fl-case-study-backend/internal/users"
	"github.com/bugrayaktiyol/fl-case-study-backend/internal/users/mocks"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_userService_ListUsers(t *testing.T) {
	t.Run("should return a list of users when valid requested", func(t *testing.T) {
		qt := assert.New(t)
		// Given
		expectedUsers := []*users.User{
			{
				Model: gorm.Model{
					ID: 1,
				},
				Name:    "John Doe",
				Age:     30,
				Address: "New York",
				Tags:    "tag1,tag2",
			},
		}

		repository := mocks.NewUserRepository(t)
		repository.On("FindAll").Return(expectedUsers, nil)
		service := users.NewService(repository)

		// When
		users, err := service.ListUsers()

		// Then
		qt.NoError(err)
		qt.Equal(expectedUsers, users)

	})
}

func Test_userService_GetUser(t *testing.T) {
	t.Run("should return a user with id when valid requested", func(t *testing.T) {
		qt := assert.New(t)
		// Given
		expectedUser := &users.User{
			Model: gorm.Model{
				ID: 1,
			},
			Name:    "John Doe",
			Age:     30,
			Address: "New York",
			Tags:    "tag1,tag2",
		}

		repository := mocks.NewUserRepository(t)
		repository.On("FindByID", "1").Return(*expectedUser, nil)
		service := users.NewService(repository)

		// When
		user, err := service.GetUser("1")

		// Then
		qt.NoError(err)
		qt.Equal(expectedUser, user)

	})

	t.Run("should return an error when invalid requested", func(t *testing.T) {
		qt := assert.New(t)
		// Given
		repository := mocks.NewUserRepository(t)
		repository.On("FindByID", "1").Return(users.User{}, errors.New("record not found"))
		service := users.NewService(repository)

		// When
		user, err := service.GetUser("1")

		// Then
		qt.Error(err)
		qt.Nil(user)

	})

}

func Test_userService_CreateUser(t *testing.T) {
	t.Run("should return a user when valid requested", func(t *testing.T) {
		qt := assert.New(t)
		// Given
		expectedUser := &users.User{
			Model: gorm.Model{
				ID: 1,
			},
			Name:    "John Doe",
			Age:     30,
			Address: "New York",
			Tags:    "tag1,tag2",
		}

		repository := mocks.NewUserRepository(t)
		repository.On("Create", expectedUser).Return(nil)
		service := users.NewService(repository)

		// When
		user, err := service.CreateUser(expectedUser)

		// Then
		qt.NoError(err)
		qt.Equal(expectedUser, user)

	})

	t.Run("should return an error when invalid requested", func(t *testing.T) {
		qt := assert.New(t)
		// Given
		expectedUser := &users.User{
			Model: gorm.Model{
				ID: 1,
			},
			Name:    "John Doe",
			Age:     30,
			Address: "New York",
			Tags:    "tag1,tag2",
		}

		repository := mocks.NewUserRepository(t)
		repository.On("Create", expectedUser).Return(errors.New("record not found"))
		service := users.NewService(repository)

		// When
		user, err := service.CreateUser(expectedUser)

		// Then
		qt.Error(err)
		qt.Nil(user)

	})

}

func Test_userService_UpdateUser(t *testing.T) {
	t.Run("should return a user when valid requested", func(t *testing.T) {
		qt := assert.New(t)
		// Given
		expectedUser := &users.User{
			Model: gorm.Model{
				ID: 1,
			},
			Name:    "John Doe",
			Age:     30,
			Address: "New York",
			Tags:    "tag1,tag2",
		}

		repository := mocks.NewUserRepository(t)
		repository.On("UpdateByID", "1", expectedUser).Return(nil)
		service := users.NewService(repository)

		// When
		user, err := service.UpdateUser("1", expectedUser)

		// Then
		qt.NoError(err)
		qt.Equal(expectedUser, user)

	})

	t.Run("should return an error when invalid requested", func(t *testing.T) {
		qt := assert.New(t)
		// Given
		expectedUser := &users.User{
			Model: gorm.Model{
				ID: 1,
			},
			Name:    "John Doe",
			Age:     30,
			Address: "New York",
			Tags:    "tag1,tag2",
		}

		repository := mocks.NewUserRepository(t)
		repository.On("UpdateByID", "1", expectedUser).Return(errors.New("record not found"))
		service := users.NewService(repository)

		// When
		user, err := service.UpdateUser("1", expectedUser)

		// Then
		qt.Error(err)
		qt.Nil(user)

	})

}
