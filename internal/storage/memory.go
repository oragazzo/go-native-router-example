package storage

import (
	"errors"

	"github.com/oragazzo/go-native-router-example/internal/models"
)

type MemoryStorage struct {
	users []models.User
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		users: make([]models.User, 0),
	}
}

func (s *MemoryStorage) GetUsers() []models.User {
	return s.users
}

func (s *MemoryStorage) AddUser(user models.User) models.User {
	s.users = append(s.users, user)
	return user
}

func (s *MemoryStorage) GetUser(id string) (models.User, error) {
	for _, user := range s.users {
		if user.ID == id {
			return user, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

func (s *MemoryStorage) UpdateUser(id string, updatedUser models.User) (models.User, error) {
	for i := range s.users {
		if s.users[i].ID == id {
			s.users[i].Name = updatedUser.Name
			return s.users[i], nil
		}
	}
	return models.User{}, errors.New("user not found")
}

func (s *MemoryStorage) DeleteUser(id string) error {
	for i := range s.users {
		if s.users[i].ID == id {
			s.users = append(s.users[:i], s.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}
