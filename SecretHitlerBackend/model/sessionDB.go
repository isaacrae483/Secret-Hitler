package model

import (
	"database/sql"
	"errors"
	"log"
)

func (session *Session) Create(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO sessions (created_at, user_id, key) VALUES (?,?,?);")
	if err != nil {
		log.Fatal(err)
	}

	result, err := stmt.Exec(session.CreatedAt, session.UserID, session.Key)
	userID, err := result.LastInsertId()
	session.ID = uint(userID)
	if err != nil {
		return err
	}

	return nil
}

func GetSession(username string, db *sql.DB) (Session, error) {
	stmt, err := db.Prepare("SELECT * FROM sessions WHERE key = ?")
	if err != nil {
		log.Fatal(err)
	}

	var session Session

	// Execute the prepared statement, passing in an id value for the
	// parameter whose placeholder is ?
	err = stmt.QueryRow(username).Scan(&session.ID, &session.CreatedAt, &session.UserID, &session.Key)
	if err != nil {
		if err == sql.ErrNoRows {
			return Session{}, errors.New("record not found")
		}
		return session, err
	}
	return session, nil
}
