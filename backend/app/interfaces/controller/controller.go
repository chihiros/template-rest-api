package controller

import (
	"net/http"
	"tamaribacms/ent"
)

type Controller struct {
	DBConn *ent.Client
}

func NewController(conn *ent.Client) *Controller {
	return &Controller{
		DBConn: conn,
	}
}

func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get, Hello, World!"))
}

func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Post, Hello, World!"))
}
