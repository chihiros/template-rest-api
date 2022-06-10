package repository

import (
	"context"
	"fmt"
	"tamaribacms/ent"
	"tamaribacms/usecase"
	"time"
)

type UserRepository struct {
	DBConn *ent.Client
}

func NewUserRepository(conn *ent.Client) *UserRepository {
	return &UserRepository{
		DBConn: conn,
	}
}

func (r *UserRepository) Get(ctx context.Context) (usecase.Response, error) {
	users, err := r.DBConn.User.Query().All(ctx)
	if err != nil {
		panic(err)
	}

	res := usecase.Response{Data: users}
	return res, err
}

func (r *UserRepository) Post(ctx context.Context, req usecase.Request) (usecase.Response, error) {
	user, err := r.DBConn.User.Create().
		SetUsername(req.Username).
		SetAge(req.Age).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		if ent.IsConstraintError(err) {
			// ent側の制約エラー
			return usecase.Response{}, fmt.Errorf("duplicate")
		}
	}

	if err != nil {
		panic(err)
	}

	res := usecase.Response{Data: user}
	return res, err
}
