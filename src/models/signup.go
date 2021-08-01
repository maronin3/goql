package models

import (
	"context"
	"goql/src/config"
)

type User struct {
	Name     string
	Password string
}

func SignUp(user User) (interface{}, error) {
	coll := config.DB.Database("User").Collection("signUp")
	result, err := coll.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}