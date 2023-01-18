package user

import (
	"context"
	"github.com/adiputraaa/bloggie/pkg/entities"
)

func (i *Interactor) FetchUsers(ctx context.Context) (entities.Users, error) {
	return i.UserRepository.FindAllUsers(ctx)
}
