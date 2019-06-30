package traph_go

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (*User, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateUser(ctx context.Context, input EditedUser) (*User, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteUser(ctx context.Context, input DeleteUser) (*MutationResult, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateProject(ctx context.Context, input NewProject) (*Project, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateProject(ctx context.Context, input EditedProject) (*Project, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateIssue(ctx context.Context, input NewIssue) (*Issue, error) {
	panic("not implemented")
}
func (r *mutationResolver) CloseIssue(ctx context.Context, input CloseIssue) (*Issue, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateComment(ctx context.Context, input NewComment) ([]*Comment, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateComment(ctx context.Context, input NewComment) ([]*Comment, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Projects(ctx context.Context) ([]*Project, error) {
	panic("not implemented")
}
func (r *queryResolver) Project(ctx context.Context, id string) (*Project, error) {
	panic("not implemented")
}
