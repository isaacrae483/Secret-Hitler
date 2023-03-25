package model

import (
	"database/sql"
	"errors"
	"log"
)

func (user *User) Create(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO users (created_at, username, password) VALUES (?,?,?);")
	if err != nil {
		log.Fatal(err)
	}

	result, err := stmt.Exec(user.CreatedAt, user.Username, user.Password)
	if err != nil {
		return err
	}
	userID, err := result.LastInsertId()
	user.ID = uint(userID)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByUsername(username string, db *sql.DB) (User, error) {
	stmt, err := db.Prepare("SELECT * FROM users WHERE username = ?")
	if err != nil {
		log.Fatal(err)
	}

	var user User

	// Execute the prepared statement, passing in an id value for the
	// parameter whose placeholder is ?
	err = stmt.QueryRow(username).Scan(&user.ID, &user.CreatedAt, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, errors.New("record not found")
		}
		return user, err
	}
	return user, nil
}

func GetUserByID(id uint, db *sql.DB) (User, error) {
	stmt, err := db.Prepare("SELECT * FROM users WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	var user User

	// Execute the prepared statement, passing in an id value for the
	// parameter whose placeholder is ?
	err = stmt.QueryRow(id).Scan(&user.ID, &user.CreatedAt, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, errors.New("record not found")
		}
		return user, err
	}
	return user, nil
}
