package user

import (
	"database/sql"
	"fmt"
	"gallery/backend/internal/types"
	"strings"
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

	err := u.db.QueryRow("SELECT id, name, email, password, avatar FROM users WHERE email = $1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Avatar)
	return user, err
}

func (u *UserRepository) DbCallUpdateUser(id string, updatedData *types.UserUpdate) (error) {
	query := "UPDATE users SET"
	args := []interface{} {}

	if updatedData.Name != nil {
		query += fmt.Sprintf("name = $%d, ", 1)
		args = append(args, *updatedData.Name)
	}

	if updatedData.Password != nil {
		query += fmt.Sprintf("password = $%d, ", 2)
		args = append(args, *updatedData.Password)
	}

	if updatedData.Avatar != nil {
		query += fmt.Sprintf("avatar = $%d, ", 3)
		args = append(args, *updatedData.Avatar)
	}

	query = strings.TrimSuffix(query, ", ")
	query += fmt.Sprintf(" WHERE id = $%d", i)
	args = append(args, id)

	_, err := u.db.Exec(query, args...)
	return err

}