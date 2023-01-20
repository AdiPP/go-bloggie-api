package user

import (
	"context"
	"github.com/adiputraaa/bloggie/pkg/entities"
	"github.com/adiputraaa/bloggie/pkg/repository"
)

type GetUserParam struct {
	User int64
}

func (i *Interactor) GetUser(ctx context.Context, param *GetUserParam) (entities.User, error) {
	return i.UserRepository.FindUser(ctx, &repository.FindUserParam{Id: param.User})
}
