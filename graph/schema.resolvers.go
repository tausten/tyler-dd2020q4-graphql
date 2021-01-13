package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/tausten/tyler-dd2020q4-graphql/graph/generated"
	"github.com/tausten/tyler-dd2020q4-graphql/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}

	r.todos = append(r.todos, todo)

	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Farm(ctx context.Context, id string) (*model.QueryFarmResult, error) {
	farmData := r.farmRepository.GetFarm(id)

	farmResult := model.QueryFarmResult{}

	if farmData != nil {
		farmResult.Data = &model.Farm{
			ID:   farmData.ID,
			Name: farmData.Name,
		}
	} else {
		farmResult.Errors = []string{"FARM NOT FOUND"}
	}

	return &farmResult, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
