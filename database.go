package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type AuthDB struct {
	connection *sql.DB
	path       string
}

func NewAuthDB(path string) *AuthDB {
	return &AuthDB{
		path: path,
	}
}

func (db *AuthDB) openConnection() error {
	conn, err := sql.Open("sqlite3", db.path)
	if err != nil {
		return err
	}

	db.connection = conn

	return nil
}

func (db *AuthDB) closeConnection() error {
	return db.connection.Close()
}

func (db *AuthDB) GetUserById(userId string) (*UserData, error) {
	err := db.openConnection()
	if err != nil {
		return nil, err
	}
	defer db.closeConnection()

	var user UserData = UserData{
		Id: userId,
	}
	query := fmt.Sprintf("SELECT username FROM users WHERE id = '%s';", userId)

	err = db.connection.QueryRow(query).Scan(&user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (db *AuthDB) GetUserByLogin(username string) (*UserData, error) {
	err := db.openConnection()
	if err != nil {
		return nil, err
	}
	defer db.closeConnection()

	var user UserData
	query := fmt.Sprintf("SELECT * FROM users WHERE username = '%s';", username)

	err = db.connection.QueryRow(query).Scan(&user.Username, &user.Password, &user.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (db *AuthDB) IsUserIdValid(userId string) (bool, error) {
	err := db.openConnection()
	if err != nil {
		return false, err
	}
	defer db.closeConnection()

	query := fmt.Sprintf("SELECT COUNT(*) FROM users WHERE id = '%s';", userId)
	var count int

	err = db.connection.QueryRow(query).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (db *AuthDB) IsPasswordCorrect(user *UserData) (bool, error) {
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

func (db *AuthDB) InsertRefteshToken(tokenId string, userId, token string) error {
	err := db.openConnection()
	if err != nil {
		return err
	}
	defer db.closeConnection()

	query := fmt.Sprintf("INSERT INTO refresh_tokens (token_id, user_id, token) VALUES ('%s', '%s', '%s');", tokenId, userId, token)
	_, err = db.connection.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (db *AuthDB) GetRefreshToken(tokenId string) (string, error) {
	err := db.openConnection()
	if err != nil {
		return "", err
	}
	defer db.closeConnection()

	query := fmt.Sprintf("SELECT token FROM refresh_tokens WHERE token_id = '%s'", tokenId)

	var token string
	err = db.connection.QueryRow(query).Scan(&token)
	if err != nil {
		return "", err
	}

	return token, nil
}
