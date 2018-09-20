package main

import (
	"./schema"
	gqlhandler "github.com/graphql-go/handler"
	"log"
	"net/http"
)

func main() {

	handler := gqlhandler.New(&gqlhandler.Config{
		Schema: &schema.Schema,
	})
	http.Handle("/graphql", handler)
	log.Println("Server started at http://localhost:8000/graphql")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
