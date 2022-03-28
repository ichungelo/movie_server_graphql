package reviews

import (
	"fmt"
	"log"
	"movie_graphql_be/graph/model"
	"movie_graphql_be/pkg/db/mysql"
	"time"
)

type Review struct {
	ReviewID string
	UserID   string
	MovieID  string
	Review   string
}

func GetAllReviewsByID(movieID string) ([]model.Review, error) {
	state, err := mysql.Db.Prepare("SELECT reviews.id, reviews.movie_id, reviews.user_id, users.username, reviews.review, reviews.created_at, reviews.updated_at FROM reviews INNER JOIN users ON users.id=reviews.user_id WHERE movie_id = ? AND is_delete = false")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rows, err := state.Query(movieID)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		return nil, err
	}

	return reviews, nil
}

func CreateReviewByID(review Review) (string, error) {
	state, err := mysql.Db.Prepare("INSERT INTO reviews (user_id, movie_id, review) VALUES (?, ?, ?)")
	if err != nil {
		log.Println(err)
		return "", err
	}

	affect, err := state.Exec(review.UserID, review.MovieID, review.Review)
	if err != nil {
		log.Println(err)
		return "", err
	}

	result, err := affect.RowsAffected()
	if err != nil {
		log.Println(err)
		return "", err
	}

	if result == 0 {
		log.Println(err)
		return "Error add review", fmt.Errorf("error add review")
	}


	return "Review Added", nil
}

func EditReviewByID(review Review) (string, error) {
	state, err := mysql.Db.Prepare("UPDATE reviews SET review=? WHERE id = ? AND user_id = ?")
	if err != nil {
		log.Println(err)
		return "", err
	}

	affect, err := state.Exec(review.Review, review.ReviewID, review.UserID)
	if err != nil {
		log.Println(err)
		return "", err
	}

	result, err := affect.RowsAffected()
	if err != nil {
		log.Println(err)
		return "", err
	}

	if result == 0 {
		err := fmt.Errorf("error edit review")
		log.Println(err)
		return "Error edit review", err
	}

	return "Review Editted", nil
}

func DeleteReviewByID(review Review) (string, error) {
	deleteTime := time.Now()
	state, err := mysql.Db.Prepare("UPDATE reviews SET is_delete = true, deleted_at= ? WHERE id = ? AND user_id = ?")
	if err != nil {
		log.Println(err)
		return "", err
	}

	affect, err := state.Exec(deleteTime, review.ReviewID, review.UserID)
	if err != nil {
		log.Println(err)
		return "", err
	}

	result, err := affect.RowsAffected()
	if err != nil {
		log.Println(err)
		return "", err
	}

	if result == 0 {
		err := fmt.Errorf("error delete review")
		log.Println(err)
		return "Error delete review", err
	}

	return "Review Deleted", nil
}
