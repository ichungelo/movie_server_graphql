package reviews

import (
	"movie_graphql_be/graph/model"
	mysql "movie_graphql_be/pkg/db"
)

type CreateReview struct {
	UserID string
	MovieID string
	Review string
}

func GetAllReviewsByID(movieID string) ([]model.Review, error) {
	state, err := mysql.Db.Prepare("SELECT reviews.id, reviews.movie_id, reviews.user_id, users.username, reviews.review, reviews.created_at, reviews.updated_at FROM reviews INNER JOIN users ON users.id=reviews.user_id WHERE movie_id = ? AND is_delete = false")
	if err != nil {
		return nil, err
	}

	rows, err := state.Query(movieID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var reviews []model.Review
	if rows.Next() {
		var review model.Review
		err := rows.Scan(&review.ID, &review.MovieID, &review.UserID, &review.Username, &review.Review, &review.CreatedAt, &review.UpdatedAt)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reviews, nil
}

func CreateReviewByID(review CreateReview) (string, error) {
	state, err := mysql.Db.Prepare("INSERT INTO reviews (user_id, movie_id, review) VALUES (?, ?, ?)")
	if err != nil {
		return "", err
	}

	_, err = state.Exec(review.UserID, review.MovieID, review.Review)
	if err != nil {
		return "", err
	}

	return "Review Added", nil
}