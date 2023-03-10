package model

import (
	"SecretHitlerBackend/utils"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type User struct {
	ID        uint
	CreatedAt time.Time
	Username  string
	Password  string
}

func (user *User) GenerateSession(db *sql.DB) (Session, error) {

	session := Session{
		CreatedAt: time.Now(),
		UserID:    user.ID,
		Key:       utils.RandString(25),
	}

	err := session.Create(db)
	if err != nil {
		return Session{}, err
	}

	return session, nil
}

func GetUserFromContext(c *gin.Context, db *sql.DB) (User, error) {
	temp, exists := c.Get("user_id")
	if !exists {
		return User{}, errors.New("no user_id")
	}
	userID := temp.(uint)
	fmt.Println("id", userID)
	return GetUserByID(userID, db)

}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (credentials *Credentials) Login(db *sql.DB) (Session, error) {
	user, err := GetUserByUsername(credentials.Username, db)
	if err != nil {
		return Session{}, err
	}

	equal := ComparePasswords(user.Password, credentials.Password)
	if !equal {
		return Session{}, errors.New("password mis match")
	}

	return user.GenerateSession(db)
}

func (credentials *Credentials) Signup(db *sql.DB) (Session, error) {
	user := User{
		CreatedAt: time.Now(),
		Username:  credentials.Username,
		Password:  HashAndSalt(credentials.Password),
	}

	if err := user.Create(db); err != nil {
		return Session{}, err
	}

	return user.GenerateSession(db)
}

// https://pkg.go.dev/golang.org/x/crypto/bcrypt
func HashAndSalt(password string) string {
	byteHash := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(byteHash, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func ComparePasswords(hashedPwd, plainPassword string) bool {
	byteHash := []byte(hashedPwd)
	bytePlain := []byte(plainPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
