package user

import (
	"ecommerce-api/domain"
	userHandler "ecommerce-api/rest/handlers/user"
)

type Service interface {
	userHandler.Service
}

type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
	// List() ([]*User, error)
	// Delete(userID int) error
	// Update(user User) (*User, error)
}
