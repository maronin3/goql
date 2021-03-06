package api

import (
	"goql/src/models"

	"github.com/graphql-go/graphql"
)

func SampleMutateType() *graphql.Object {
	sampleType := graphql.NewObject(graphql.ObjectConfig{
		Name: "SampleMutate",
		Fields: graphql.Fields{
			"Message": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
	return sampleType
}

func SampleMutateArgs() map[string]*graphql.ArgumentConfig {
	sampleArgs := graphql.FieldConfigArgument{
		"Username": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"Age": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
	return sampleArgs
}

func SampleMutateResolve(params graphql.ResolveParams) (interface{}, error) {
	resp := &models.SampleMutateResp{
		Message: "Hello, This is SampleMutate",
	}
	return resp, nil
}
