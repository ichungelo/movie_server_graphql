package movies

import (
	"log"
	"movie_graphql_be/graph/model"
	"movie_graphql_be/pkg/db/mysql"
)

func GetAll() ([]model.Movie, error) {
	state, err := mysql.Db.Prepare("SELECT id, title, release_year, production, overview FROM movies")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rows, err := state.Query()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	var movies []model.Movie

	for rows.Next() {
		var movie model.Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Year, &movie.Poster, &movie.Overview)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return movies, nil
}

func GetByID(id string) (model.Movie, error) {
	state, err := mysql.Db.Prepare("SELECT id, title, release_year, production, overview FROM movies WHERE id = ?")
	if err != nil {
		log.Println(err)
		return model.Movie{}, err
	}

	rows, err := state.Query(id)
	if err != nil {
		log.Println(err)
		return model.Movie{}, err
	}

	defer rows.Close()

	var movie model.Movie
	if rows.Next() {
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Year, &movie.Poster, &movie.Overview)
		if err != nil {
			log.Println(err)
			return model.Movie{}, err
		}
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return model.Movie{}, err
	}

	return movie, nil
}
