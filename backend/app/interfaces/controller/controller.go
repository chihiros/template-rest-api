package controller

import (
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
