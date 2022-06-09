package controller

import (
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
	w.Write([]byte("Post, Hello, World!"))
}
