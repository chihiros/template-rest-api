package repository

import (
	"context"
	"tamaribacms/ent"
	"tamaribacms/entity"
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

	}

	return us, err
}
