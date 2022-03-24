package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"movie_graphql_be/graph/generated"
	"movie_graphql_be/graph/model"
	"movie_graphql_be/internal/users"
)

func (r *mutationResolver) DetailMovie(ctx context.Context, input model.PrimaryID) (*model.MovieDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Register(ctx context.Context, input model.Register) (string, error) {
	var user users.User

	user.Username = input.Username
	user.Email = input.Email
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Password = input.Password
	user.ConfirmPassword = input.ConfirmPassword
	success, err := user.CreateUser()
	if err != nil {
		return "", err
	}

	return success, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var login users.Login
	login.Username = input.Username
	login.Password = input.Password

	token, err := login.LoginUser()
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *mutationResolver) NewReview(ctx context.Context, input model.NewReview) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Movies(ctx context.Context) ([]*model.Movie, error) {
	var movies []*model.Movie
	dummy := model.Movie{
		ID:       "1",
		Title:    "Dummy Movie",
		Year:     2020,
		Poster:   "https://m.media-amazon.com/images/M/MV5BMjIxMDgxMzc4MV5BMl5BanBnXkFtZTgwMzQzMzMzMjE@._V1_SX300.jpg",
		Overview: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
	}

	movies = append(movies, &dummy)
	return movies, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
