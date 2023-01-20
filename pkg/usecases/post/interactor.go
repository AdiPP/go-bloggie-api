package post

import "github.com/adiputraaa/bloggie/pkg/repository"

type Interactor struct {
	PostRepository repository.IPostRepository
}

func NewInteractor(postRepository repository.IPostRepository) *Interactor {
	return &Interactor{PostRepository: postRepository}
}
