package api

import (
	api "goql/src/api/sample"

	"github.com/graphql-go/graphql"
)

func SetMutation() graphql.Fields {
	mutateQuery := graphql.Fields{
		"SampleMutate": &graphql.Field{
			Type:    api.SampleMutateType(),
			Args:    api.SampleMutateArgs(),
			Resolve: api.SampleMutateResolve,
		},
	}
	return mutateQuery
}
