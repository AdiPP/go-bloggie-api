package post

type Interactor struct {
	PostRepository IPostRepository
}

func NewInteractor(postRepository IPostRepository) *Interactor {
	return &Interactor{PostRepository: postRepository}
}
