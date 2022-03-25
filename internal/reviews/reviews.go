package reviews

import (
	"movie_graphql_be/graph/model"
	mysql "movie_graphql_be/pkg/db"
)

func GetAllReviewsByID(movieID string) ([]model.Review, error) {
	state, err := mysql.Db.Prepare("SELECT reviews.id, reviews.movie_id, reviews.user_id, users.username, reviews.review, reviews.created_at, reviews.updated_at FROM reviews INNER JOIN users ON users.id=reviews.movie_id WHERE movie_id = ? AND is_delete = false")
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