package api

import (
	api "goql/src/api/sample"

	"github.com/graphql-go/graphql"
)

func SetQuery() graphql.Fields {
	query := graphql.Fields{
		"SampleQuery": &graphql.Field{
			Type:    api.SampleQueryType(),
			Args:    api.SampleQueryArgs(),
			Resolve: api.SampleQueryResolve,
		},
	}
	return query
}
