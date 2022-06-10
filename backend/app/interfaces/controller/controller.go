package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"tamaribacms/ent"
	"tamaribacms/interfaces/repository"
	"tamaribacms/usecase"
)

type Controller struct {
	Usecase usecase.UserUseCase
}

func NewController(conn *ent.Client) *Controller {
	u := NewUserUsecase(conn)
	return &Controller{
		Usecase: u,
	}
}

func NewUserUsecase(conn *ent.Client) *usecase.UserUsecase {
	repo := repository.NewUserRepository(conn)
	return &usecase.UserUsecase{
		Repository: repo,
	}
}

func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	users, err := c.Usecase.Get(context.Background())
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
	// bodyの中身をbindする
	req := usecase.Request{}
	err := json.NewDecoder(r.Body).Decode(&req)
	user, err := c.Usecase.Post(context.Background(), req)

	if err != nil {
		switch err.Error() {
		case "duplicate":
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(user)
		default:
			panic(err)
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
