package repository

import (
	"context"
	"database/sql"

	"github.com/ropel12/simple-ecommerce-api/model"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user model.UserRegis) error
	Update(ctx context.Context, tx *sql.Tx, user model.UserUpdate, id int) error
	FindById(ctx context.Context, tx *sql.Tx, userId int) (*model.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]*model.UserAll, error)
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (*model.User, error)
}
