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

func (r *UserRepository) Get(ctx context.Context) ([]entity.User, error) {
	users, err := r.DBConn.User.Query().All(ctx)
	if err != nil {
		panic(err)
	}

	us := []entity.User{}
	for _, user := range users {
		us = append(us, entity.User{
			ID:        user.ID,
			Username:  user.Username,
			Age:       user.Age,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return us, err
}
