package api

import (
	"goql/src/models"

	"github.com/graphql-go/graphql"
)

func SampleQueryType() *graphql.Object {
	sampleType := graphql.NewObject(graphql.ObjectConfig{
		Name: "SampleQuery",
		Fields: graphql.Fields{
			"Message": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
	return sampleType
}

func SampleQueryArgs() map[string]*graphql.ArgumentConfig {
	sampleArgs := graphql.FieldConfigArgument{
		"ID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"Password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}
	return sampleArgs
}

func SampleQueryResolve(params graphql.ResolveParams) (interface{}, error) {
	resp := &models.SampleMutateResp{
		Message: "Hello, This is SampleQuery",
	}
	return resp, nil
}
