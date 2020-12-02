package graph

import (
	"github.com/tausten/tyler-dd2020q4-graphql/graph/model"
	"github.com/tausten/tyler-dd2020q4-graphql/repository"
)

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos          []*model.Todo
	farmRepository repository.FarmRepository
}

func CreateResolver(farmRepository repository.FarmRepository) *Resolver {
	return &Resolver{
		farmRepository: farmRepository,
	}
}
