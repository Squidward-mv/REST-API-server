package store

import "RestAPI/internal/app/model"

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(string) (*model.User, error)
}
