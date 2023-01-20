package repository

import (
	"context"
	"github.com/adiputraaa/bloggie/pkg/entities"
)

type (
	FindAllPostsByUserParam struct {
		User int64
	}

	FindUserParam struct {
		Id int64
	}
)

type IPostRepository interface {
	FindAllPosts(ctx context.Context) (posts entities.Posts, err error)
	FindAllPostsByUser(ctx context.Context, param *FindAllPostsByUserParam) (posts entities.Posts, err error)
}

type IUserRepository interface {
	FindUser(ctx context.Context, param *FindUserParam) (user entities.User, err error)
	FindAllUsers(ctx context.Context) (users entities.Users, err error)
}
