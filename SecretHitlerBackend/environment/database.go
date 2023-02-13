package environment

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func Connect(fromLower bool) *AppConfig {
	file := "./secretHitler.db"
	if fromLower {
		file = "." + file
	}
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(1)
	//defer func(db *sql.DB) {
	//	err := db.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}(db)

	return &AppConfig{
		DB: db,
	}
}
