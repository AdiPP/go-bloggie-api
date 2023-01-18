package post

import (
	"context"
	"github.com/adiputraaa/bloggie/pkg/entities"
)

type IPostRepository interface {
	FindAllPosts(ctx context.Context) (posts entities.Posts, err error)
	FindAllPostsByUser(ctx context.Context, param *FindAllPostsByUserParam) (posts entities.Posts, err error)
}

type (
	FindAllPostsByUserParam struct {
		User int64
	}
)
