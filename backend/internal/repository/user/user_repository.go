package user

import (
	"database/sql"
	"gallery/backend/internal/types"
)

type UserRepositoryInterface interface {
	DbCallCreateUser(name, email, password, avatar string) (int, error)
	DbCallGetUserByEmail(email string) (types.UserModel, error)
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

func (u *UserRepository) DbCallGetUserByEmail(email string) (types.UserModel, error) {
	var user types.UserModel

	err := u.db.QueryRow("SELECT FROM users (id, name, email, password, avatar) WHERE email = $1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Avatar)
	return user, err
}