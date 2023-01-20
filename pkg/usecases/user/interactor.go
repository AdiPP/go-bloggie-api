package user

import "github.com/adiputraaa/bloggie/pkg/repository"

type Interactor struct {
	UserRepository repository.IUserRepository
}

func NewInteractor(repository repository.IUserRepository) *Interactor {
	return &Interactor{UserRepository: repository}
}
