package repository

import (
	"context"
	"tamaribacms/ent"
)

type UserRepository struct {
	DBConn *ent.Client
}

func NewUserRepository(conn *ent.Client) *UserRepository {
	return &UserRepository{
		DBConn: conn,
	}
}

func (r *UserRepository) Get(ctx context.Context) string {
	return "Get, Hello, World!"
}
