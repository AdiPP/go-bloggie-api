package post

import (
	"context"
	"github.com/adiputraaa/bloggie/pkg/entities"
)

func (i *Interactor) FetchPosts(ctx context.Context) (entities.Posts, error) {
	return i.PostRepository.FindAllPosts(ctx)
}
