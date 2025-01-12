package user

import (
	"errors"
	"sync"

	"github.com/ithinkiborkedit/GUSH.git/pkg/models"
)

type UserManager struct {
	users     map[string]*models.User
	usersLock sync.RWMutex
}

func NewUserManager() *UserManager {
	return &UserManager{
		users: make(map[string]*models.User),
	}
}

func (um *UserManager) Authenticate(username string) (*models.User, error) {
	um.usersLock.Lock()
	defer um.usersLock.Unlock()

	if username == "" {
		return nil, errors.New("username cannot be empty")
	}
	if usr, exists := um.users[username]; exists {
		return usr, nil
	}

	usr := &models.User{
		Username: username,
	}

	return usr, nil
}
