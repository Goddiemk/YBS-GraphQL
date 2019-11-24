package main

import (
	"minkhantkoko/YBS/lib"
	"net/http"

	"github.com/graphql-go/handler"
)

func main() {
	schema, _ := lib.BaseSchema()
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
