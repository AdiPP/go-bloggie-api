package user

import (
	"context"
	"github.com/adiputraaa/bloggie/pkg/entities"
)

type GetUserParam struct {
	User int64
}

func (i *Interactor) GetUser(ctx context.Context, param *GetUserParam) (entities.User, error) {
	return i.UserRepository.FindUser(ctx, &FindUserParam{Id: param.User})
}
