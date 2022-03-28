package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"

	"fmt"
	"os"
)

var Db *sql.DB

func InitDB() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?query", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err.Error())
	}

	if err := db.Ping(); err != nil {
		panic(err.Error())
	}

	log.Println("Successfully connected to MySQL")
	Db = db
}

func Migrate() {
	if err := Db.Ping(); err != nil {
		panic(err.Error())
	}

	driver, _ := mysql.WithInstance(Db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://pkg/db/migration/mysql",
		"mysql",
		driver,
	)

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err.Error())
	}
}
