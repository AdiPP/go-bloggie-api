package sqlite

import (
	"context"
	"github.com/adiputraaa/bloggie/pkg/entities"
	"github.com/adiputraaa/bloggie/pkg/usecases/user"
	"github.com/doug-martin/goqu/v9"
)

type Repository struct {
	Database *goqu.Database
}

func NewRepository(database *goqu.Database) *Repository {
	return &Repository{Database: database}
}

func (ur *Repository) FindUser(ctx context.Context, param *user.FindUserParam) (user entities.User, err error) {
	dataset := ur.Database.From("users").Select("id", "username", "password", "email")
	dataset = dataset.Where(goqu.Ex{"id": param.Id})

	_, err = dataset.Executor().ScanStructContext(ctx, &user)

	return
}

func (ur *Repository) FindAllUsers(ctx context.Context) (users entities.Users, err error) {
	dataset := ur.Database.From("users").Select("id", "username", "password", "email")

	err = dataset.Executor().ScanStructsContext(ctx, &users)

	return
}
