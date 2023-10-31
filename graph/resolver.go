package graph

import "github.com/Purvig648/graphql-demo/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen

type Resolver struct{
	Service service.UserService
}
