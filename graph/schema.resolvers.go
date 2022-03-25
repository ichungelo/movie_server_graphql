package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"movie_graphql_be/graph/generated"
	"movie_graphql_be/graph/model"
	"movie_graphql_be/internal/auth"
	"movie_graphql_be/internal/movies"
	"movie_graphql_be/internal/reviews"
	"movie_graphql_be/internal/users"
)

func (r *mutationResolver) DetailMovie(ctx context.Context, input model.PrimaryID) (*model.MovieDetail, error) {
	var resultMovie *model.MovieDetail
	var resultReviews []*model.Review
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("not authorized")
	}

	dbMovie, err := movies.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	dbReviews, err := reviews.GetAllReviewsByID(input.ID)
	if err != nil {
		return nil, err
	}

	for _, dbReviews := range dbReviews {
		resultReviews = append(resultReviews, &model.Review{
			ID:       dbReviews.ID,
			MovieID:  dbReviews.MovieID,
			UserID:   dbReviews.UserID,
			Username: dbReviews.Username,
			Review:   dbReviews.Review,
			CreatedAt: dbReviews.CreatedAt,
			UpdatedAt: dbReviews.UpdatedAt,
		})
	}

	resultMovie = &model.MovieDetail{
		ID:       dbMovie.ID,
		Title:    dbMovie.Title,
		Year:     dbMovie.Year,
		Poster:   dbMovie.Poster,
		Overview: dbMovie.Overview,
		Reviews: resultReviews,
	}

	return resultMovie, nil
}

func (r *mutationResolver) Register(ctx context.Context, input model.Register) (string, error) {
	var user users.User
	var login users.Login

	user.Username = input.Username
	user.Email = input.Email
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Password = input.Password
	user.ConfirmPassword = input.ConfirmPassword
	_, err := user.CreateUser()
	if err != nil {
		return "", err
	}

	login.Username = user.Username
	login.Password = user.Password
	token, err := login.LoginUser()
	if err != nil {
		return "", err
	}

	bearerToken := "Bearer " + token

	return bearerToken, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var login users.Login
	login.Username = input.Username
	login.Password = input.Password

	token, err := login.LoginUser()
	if err != nil {
		return "", err
	}

	berarerToken := "Bearer " + token

	return berarerToken, nil
}

func (r *mutationResolver) NewReview(ctx context.Context, input model.NewReview) (string, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return "", fmt.Errorf("not authorized")
	}
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Movies(ctx context.Context) ([]*model.Movie, error) {
	var result []*model.Movie
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("not authorized")
	}

	dbMovies, err := movies.GetAll()
	if err != nil {
		return nil, err
	}

	for _, dbMovie := range dbMovies {
		result = append(result, &model.Movie{
			ID:       dbMovie.ID,
			Title:    dbMovie.Title,
			Year:     dbMovie.Year,
			Poster:   dbMovie.Poster,
			Overview: dbMovie.Overview,
		})
	}

	return result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
