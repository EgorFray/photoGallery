package user

import (
	"database/sql"
)

type UserInterface interface {
	DbCallCreateUser(name, email, password string) (int, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) DbCallCreateUser(name, email, password, avatar string) (int, error) {
	var insertedID int

	err := u.db.QueryRow("INSERT INTO users (name, email, password, avatar) VALUES ($1, $2, $3, $4) RETURNING id", name, email, password, avatar).Scan(&insertedID)
	return insertedID, err
}