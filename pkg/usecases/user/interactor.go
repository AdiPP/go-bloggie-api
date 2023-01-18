package user

type Interactor struct {
	UserRepository IUserRepository
}

func NewInteractor(repository IUserRepository) *Interactor {
	return &Interactor{UserRepository: repository}
}
