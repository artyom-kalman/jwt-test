package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type UsersDatabase struct {
	connection *sql.DB
	path       string
}

func NewUserDatabase(path string) *UsersDatabase {
	return &UsersDatabase{
		path: path,
	}
}

func (db *UsersDatabase) openConnection() error {
	conn, err := sql.Open("sqlite3", db.path)
	if err != nil {
		return err
	}

	db.connection = conn

	return nil
}

func (db *UsersDatabase) closeConnection() error {
	return db.connection.Close()
}

func (db *UsersDatabase) UserExists(user *UserData) (bool, error) {
	err := db.openConnection()
	if err != nil {
		return false, err
	}
	defer db.closeConnection()

	query := fmt.Sprintf("SELECT COUNT(*) FROM users WHERE username = '%s';", user.Username)
	var count int

	err = db.connection.QueryRow(query).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (db *UsersDatabase) IsPasswordCorrect(user *UserData) (bool, error) {
	err := db.openConnection()
	if err != nil {
		return false, err
	}
	defer db.closeConnection()

	var password string
	query := fmt.Sprintf("SELECT password FROM users WHERE username = '%s';", user.Username)
	err = db.connection.QueryRow(query, user.Username).Scan(&password)
	if err != nil {
		return false, err
	}

	return user.Password == password, nil
}
