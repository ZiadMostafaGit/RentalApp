package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ZiadMostafaGit/rental-app/internal/models"
)

type User_repo struct {
	DB *sql.DB
}

func New_user_repo(db *sql.DB) *User_repo {
	return &User_repo{
		DB: db,
	}
}

func (r *User_repo) Get_by_id(ctx context.Context, id uint) (*models.User, error) {
	query := `SELECT id, first_name, last_name, role, gender, state, city, street, score, email 
	          FROM users WHERE id = ?`

	user := &models.User{}
	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&user.Id,
		&user.First_name,
		&user.Last_name,
		&user.Role,
		&user.Gender,
		&user.State,
		&user.City,
		&user.Street,
		&user.Score,
		&user.Email,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		return nil, err
	}
	return user, nil
}

func (r *User_repo) Insert_user(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (first_name, last_name, role, gender, state, city, street, email, password)
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := r.DB.ExecContext(ctx, query, user.First_name, user.Last_name, user.Role, user.Gender, user.State, user.City, user.Street, user.Email, user.Password)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	if id != int64(user.Id) {
		return fmt.Errorf("error not inserted")
	}
	return nil
}

// DELETE user by ID
func (r *User_repo) Delete_user(ctx context.Context, id uint) error {
	query := `DELETE FROM users WHERE id = ?`
	result, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id %d", id)
	}
	return nil
}

// UPDATE all fields
func (r *User_repo) Update_user(ctx context.Context, user *models.User) error {
	query := `UPDATE users 
	          SET first_name = ?, last_name = ?, role = ?, gender = ?, state = ?, city = ?, street = ?, email = ?, password = ?
	          WHERE id = ?`

	result, err := r.DB.ExecContext(ctx, query, user.First_name, user.Last_name, user.Role, user.Gender, user.State, user.City, user.Street, user.Email, user.Password, user.Id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id %d to update", user.Id)
	}
	return nil
}

// INDIVIDUAL FIELD UPDATES
func (r *User_repo) Update_user_first_name(ctx context.Context, id uint, firstName string) error {
	return r.updateField(ctx, "first_name", firstName, id)
}

func (r *User_repo) Update_user_last_name(ctx context.Context, id uint, lastName string) error {
	return r.updateField(ctx, "last_name", lastName, id)
}

func (r *User_repo) Update_user_role(ctx context.Context, id uint, role string) error {
	return r.updateField(ctx, "role", role, id)
}

func (r *User_repo) Update_user_gender(ctx context.Context, id uint, gender string) error {
	return r.updateField(ctx, "gender", gender, id)
}

func (r *User_repo) Update_user_state(ctx context.Context, id uint, state string) error {
	return r.updateField(ctx, "state", state, id)
}

func (r *User_repo) Update_user_city(ctx context.Context, id uint, city string) error {
	return r.updateField(ctx, "city", city, id)
}

func (r *User_repo) Update_user_street(ctx context.Context, id uint, street string) error {
	return r.updateField(ctx, "street", street, id)
}

func (r *User_repo) Update_user_email(ctx context.Context, id uint, email string) error {
	return r.updateField(ctx, "email", email, id)
}

func (r *User_repo) Update_user_password(ctx context.Context, id uint, password string) error {
	return r.updateField(ctx, "password", password, id)
}

// helper method
func (r *User_repo) updateField(ctx context.Context, field string, value any, id uint) error {
	query := fmt.Sprintf("UPDATE users SET %s = ? WHERE id = ?", field)
	result, err := r.DB.ExecContext(ctx, query, value, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id %d to update field %s", id, field)
	}
	return nil
}
