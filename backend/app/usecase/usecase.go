package usecase

import (
	"context"
)

type UserUseCase interface {
	Get(context.Context) (Response, error)
	Post(context.Context, Request) (Response, error)
}

type UserRepository interface {
	Get(context.Context) (Response, error)
	Post(context.Context, Request) (Response, error)
}

type UserUsecase struct {
	Repository UserRepository
}

func (u *UserUsecase) Get(ctx context.Context) (Response, error) {
	return u.Repository.Get(ctx)
}

func (u *UserUsecase) Post(ctx context.Context, req Request) (Response, error) {
	return u.Repository.Post(ctx, req)
}
