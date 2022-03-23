package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"movie_graphql_be/graph/generated"
	"movie_graphql_be/graph/model"
)

func (r *mutationResolver) Movie(ctx context.Context, input model.PrimaryID) (*model.MovieDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) User(ctx context.Context, input model.PrimaryID) (*model.UserDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Register(ctx context.Context, input model.Register) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) NewReview(ctx context.Context, input model.NewReview) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Movies(ctx context.Context) ([]*model.Movie, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
