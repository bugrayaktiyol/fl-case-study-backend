package users

type UserService struct {
	repository *UserRepository
}

func NewUserService(repository *UserRepository) *UserService {
	return &UserService{repository: repository}
}

// ListUsers returns a list of all users
func (us *UserService) ListUsers() ([]User, error) {
	return us.repository.FindAll()
}

// GetUser returns a specific user by ID
func (us *UserService) GetUser(id string) (*User, error) {
	user, err := us.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user
func (us *UserService) CreateUser(user *User) (*User, error) {
	if err := us.repository.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser updates a specific user by ID
func (us *UserService) UpdateUser(id string, updatedUser *User) (*User, error) {
	if err := us.repository.UpdateByID(id, updatedUser); err != nil {
		return nil, err
	}
	return updatedUser, nil
}

// DeleteUser deletes a specific user by ID
func (us *UserService) DeleteUser(id string) error {
	return us.repository.DeleteByID(id)
}
