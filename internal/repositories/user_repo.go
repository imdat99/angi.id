package repositories

import (
	"context"
	"database/sql"

	"angi.id/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id string) error

	GetListOfUsers(ctx context.Context, filter models.UserFilter) ([]models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := "INSERT INTO users (id, name, email) VALUES (?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, user.Id, user.Name, user.Email)
	return err
}

func (r *userRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := r.db.QueryRowContext(ctx, query, id)

	user := &models.User{}
	if err := row.Scan(&user.Id, &user.Name, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *models.User) error {
	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Id)
	return err
}

func (r *userRepository) DeleteUser(ctx context.Context, id string) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *userRepository) GetListOfUsers(ctx context.Context, filter models.UserFilter) ([]models.User, error) {
	query := "SELECT id, name, email FROM users WHERE 1=1 LIMIT ? OFFSET ?"
	args := []interface{}{filter.Limit, filter.Offset}

	if filter.Email != nil {
		query += " AND email = ?"
		args = append(args, *filter.Email)
	}
	if filter.IsActive != nil {
		query += " AND is_active = ?"
		args = append(args, *filter.IsActive)
	}
	if filter.CreatedAfter != nil {
		query += " AND creation_date > ?"
		args = append(args, *filter.CreatedAfter)
	}
	if filter.CreatedBefore != nil {
		query += " AND creation_date < ?"
		args = append(args, *filter.CreatedBefore)
	}
	if filter.LastLoginAfter != nil {
		query += " AND last_login > ?"
		args = append(args, *filter.LastLoginAfter)
	}
	if filter.LastLoginBefore != nil {
		query += " AND last_login < ?"
		args = append(args, *filter.LastLoginBefore)
	}
	if filter.Name != nil {
		query += " AND name LIKE ?"
		args = append(args, "%"+*filter.Name+"%")
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
