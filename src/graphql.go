package src

import (
	"log"
	"net/http"

	"goql/src/api"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func GraphqlSetting() http.Handler {
	query := api.SetQuery()
	mutateQuery := api.SetMutation()
	rootMutation := graphql.NewObject(graphql.ObjectConfig{Name: "RootMutation", Fields: mutateQuery})
	rootQuery := graphql.NewObject(graphql.ObjectConfig{Name: "RootQuery", Fields: query})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
	if err != nil {
		log.Println(err)
	}
	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	return h
}
