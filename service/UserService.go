package service

import (
	"context"

	"github.com/ropel12/simple-ecommerce-api/model"
)

type UserService_impl interface {
	CreateUser(ctx context.Context, request model.UserRegis) error
	UpdateUser(ctx context.Context, request model.UserUpdate, id int) error
	FindAllUser(ctx context.Context) ([]*model.UserAll, error)
	FindUserById(ctx context.Context, userid int) (*model.User, error)
	Login(ctx context.Context, req model.LoginRequest) (string, error)
	FindUserByUsername(ctx context.Context, username string) (*model.User, error)
}
