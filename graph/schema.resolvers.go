package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/tausten/tyler-dd2020q4-graphql/graph/generated"
	"github.com/tausten/tyler-dd2020q4-graphql/graph/model"
	"github.com/tausten/tyler-dd2020q4-graphql/repository"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	newItem := &repository.TodoItem{
		ID:     nil,
		Text:   input.Text,
		Done:   false,
		UserID: input.UserID,
	}

	items, errors := r.todoRepository.Upsert([]*repository.TodoItem{newItem})

	// NOTE: The repository can take care of if/when it needs to ensure referential integrity of the user foreign key, and report
	// violations as error here.
	if errors[0] != nil {
		return nil, errors[0]
	}

	todo := &model.Todo{
		Text:   items[0].Text,
		ID:     items[0].ID.Value,
		UserID: items[0].UserID,
		Done:   items[0].Done,
	}

	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) (*model.TodoConnection, error) {
	items, err := r.todoRepository.GetTodos()
	if err != nil {
		return nil, err
	}

	result := &model.TodoConnection{
		PageInfo: &model.PageInfo{},
	}

	if len(items) < 1 {
		return result, nil
	}

	result.Edges = make([]*model.TodoEdge, len(items))

	for i, v := range items {
		result.Edges[i] = &model.TodoEdge{
			Node:   v.ToTodo(),
			Cursor: v.EncodeCursor(i),
		}
	}

	result.PageInfo.StartCursor = result.Edges[0].Cursor
	result.PageInfo.EndCursor = result.Edges[len(result.Edges)-1].Cursor

	return result, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	matches, errors := r.todoRepository.GetTodosById([]string{id})

	if errors[0] != nil {
		return nil, errors[0]
	}

	return matches[0].ToTodo(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
