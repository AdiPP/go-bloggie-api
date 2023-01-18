package user

import (
	"context"
	"github.com/adiputraaa/bloggie/pkg/entities"
)

type FindUserParam struct {
	Id int64
}

type IUserRepository interface {
	FindUser(ctx context.Context, param *FindUserParam) (user entities.User, err error)
	FindAllUsers(ctx context.Context) (users entities.Users, err error)
}
