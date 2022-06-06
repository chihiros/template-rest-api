package main

import (
	"net/http"
	"tamaribacms/infra"
)

func main() {
	r := infra.NewRouter()
	err := http.ListenAndServe(":9876", r)
	if err != nil {
		panic(err)
	}
}
