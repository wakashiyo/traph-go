package traph_go

import (
	"context"
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	todos []*Todo
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*Todo, error) {

	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
	d := rand.Int63()

	todo := &Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", d),
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*Todo, error) {
	return r.todos, nil
}
func (r *queryResolver) Todo(ctx context.Context, id string) (*Todo, error) {
	t := r.todos[0]
	return t, nil
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) Done(ctx context.Context, obj *Todo) (*bool, error) {
	panic("not implemented")
}
func (r *todoResolver) User(ctx context.Context, obj *Todo) (*User, error) {
	name := "user " + obj.UserID
	return &User{ID: obj.UserID, Name: &name}, nil
}
