package users

type UserService interface {
	ListUsers() ([]*User, error)
	GetUser(id string) (*User, error)
	CreateUser(user *User) (*User, error)
	UpdateUser(id string, updatedUser *User) (*User, error)
	DeleteUser(id string) error
}

type userService struct {
	repository UserRepository
}

func NewService(repository UserRepository) UserService {
	return &userService{repository: repository}
}

// ListUsers returns a list of all users
func (us *userService) ListUsers() ([]*User, error) {
	return us.repository.FindAll()
}

// GetUser returns a specific user by ID
func (us *userService) GetUser(id string) (*User, error) {
	user, err := us.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user
func (us *userService) CreateUser(user *User) (*User, error) {
	if err := us.repository.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser updates a specific user by ID
func (us *userService) UpdateUser(id string, updatedUser *User) (*User, error) {
	if err := us.repository.UpdateByID(id, updatedUser); err != nil {
		return nil, err
	}
	return updatedUser, nil
}

// DeleteUser deletes a specific user by ID
func (us *userService) DeleteUser(id string) error {
	return us.repository.DeleteByID(id)
}
