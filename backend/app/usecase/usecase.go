package usecase

import (
	"context"
	"tamaribacms/entity"
)

type UserUseCase interface {
	Get(context.Context) (Response, error)
	// Post(context.Context) (entity.User, error)
}

type UserRepository interface {
	Get(context.Context) (Response, error)
	// Post(context.Context) (entity.User, error)
}

type UserUsecase struct {
	Repository UserRepository
}

func (u *UserUsecase) Get(ctx context.Context) (Response, error) {
	return u.Repository.Get(ctx)
}
