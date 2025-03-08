package repository

import (
	"database/sql"
	"time"

	"github.com/ryota1119/gin_webapi/internal/domain"
	"github.com/ryota1119/gin_webapi/internal/domain/repository"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

var _ repository.UserRepository = (*UserRepositoryImpl)(nil)

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepositoryImpl{db}

}

func (r *UserRepositoryImpl) Create(user *domain.User) error {
	query := "INSERT INTO `users` (`name`, `email`, `password`, `created_at`, `updated_at`) " +
		"VALUES (?, ?, ?, ?, ?)"
	_, err := r.db.Exec(
		query,
		user.Name,
		user.Email,
		user.Password,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) Find(userID domain.UserID) (*domain.User, error) {
	var user domain.User

	query := "SELECT `id`, `name`, `email`, `password` FROM `users` WHERE `id` = ?"
	result := r.db.QueryRow(query, userID)
	err := result.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*domain.User, error) {
	var user domain.User

	query := "SELECT `id`, `name`, `email`, `password` FROM `users` WHERE `email` = ?"
	result := r.db.QueryRow(query, email)
	err := result.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	return &user, nil
}
