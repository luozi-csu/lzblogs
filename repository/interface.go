package repository

import (
	"context"

	"github.com/luozi-csu/lzblogs/model"
)

type Repository interface {
	User() UserRepository
	Ping(ctx context.Context) error
	Init() error
	Migrant
}

type Migrant interface {
	Migrate() error
}

type UserRepository interface {
	GetUserByID(uint) (*model.User, error)
	GetUserByAuthID(authType, authID string) (*model.User, error)
	GetUserByName(string) (*model.User, error)
	List() (model.Users, error)
	Create(*model.User) (*model.User, error)
	Update(*model.User) (*model.User, error)
	Delete(*model.User) error
	AddRole(role *model.Role, user *model.User) error
	DelRole(role *model.Role, user *model.User) error
	Migrate() error
}