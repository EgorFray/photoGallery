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
	DbCallGetUserById(id string) (types.UserResponse, error)
	DbCallUpdateUser(id string, name, password, avatar *string ) (error)
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

func (u *UserRepository) DbCallGetUserById(id string) (types.UserResponse, error) {
	var user types.UserResponse

	err := u.db.QueryRow("SELECT id, name, email, avatar FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Avatar)
	return user, err
}

func (u *UserRepository) DbCallUpdateUser(id string, name, password, avatar *string ) (error) {
	query := "UPDATE users SET "
	args := []interface{} {}
	i := 1

	if name != nil {
		query += fmt.Sprintf("name = $%d, ", i)
		args = append(args, *name)
		i++
	}

	if password != nil {
		query += fmt.Sprintf("password = $%d, ", i)
		args = append(args, *password)
		i++
	}

	if avatar != nil {
		query += fmt.Sprintf("avatar = $%d, ", i)
		args = append(args, *avatar)
		i++
	}

	if len(args) == 0 {
		return nil
}

	query = strings.TrimSuffix(query, ", ")
	query += fmt.Sprintf(" WHERE id = $%d", i)
	args = append(args, id)

	_, err := u.db.Exec(query, args...)
	return err
}