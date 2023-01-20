package post

import (
	"context"
	"github.com/adiputraaa/bloggie/pkg/entities"
	"github.com/adiputraaa/bloggie/pkg/repository"
)

type FetchPostsByUserParam struct {
	User int64
}

func (i *Interactor) FetchPostsByUser(ctx context.Context, param *FetchPostsByUserParam) (entities.Posts, error) {
	return i.PostRepository.FindAllPostsByUser(ctx, &repository.FindAllPostsByUserParam{User: param.User})
}
