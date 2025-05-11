package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ZiadMostafaGit/rental-app/internal/models"
)

type Rental_repo struct {
	DB *sql.DB
}

func New_rental_repo(db *sql.DB) *Rental_repo {
	return &Rental_repo{DB: db}
}

func (r *Rental_repo) Get_rental(ctx context.Context, id uint) (*models.Rental, error) {
	query := `SELECT id, item_id, user_id, start_date, end_date, current_states, estimated_time, delivery_address, created_at 
	          FROM rentals WHERE id = ?`

	rental := &models.Rental{}
	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&rental.Id,
		&rental.Item_id,
		&rental.User_id,
		&rental.Start_date,
		&rental.End_date,
		&rental.Current_state,
		&rental.Esstimated_time,
		&rental.Delivery_address,
		&rental.Created_at,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no rental found with id %d", id)
		}
		return nil, err
	}
	return rental, nil
}

func (r *Rental_repo) Insert_rental(ctx context.Context, rental *models.Rental) error {
	query := `INSERT INTO rentals (item_id, user_id, start_date, end_date, current_states, estimated_time, delivery_address, created_at)
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.DB.ExecContext(ctx, query,
		rental.Item_id,
		rental.User_id,
		rental.Start_date,
		rental.End_date,
		rental.Current_state,
		rental.Esstimated_time,
		rental.Delivery_address,
		rental.Created_at,
	)
	return err
}

func (r *Rental_repo) Delete_rental(ctx context.Context, id uint) error {
	query := `DELETE FROM rentals WHERE id = ?`
	result, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("no rental found with id %d", id)
	}
	return nil
}

func (r *Rental_repo) Update_rental(ctx context.Context, rental *models.Rental) error {
	query := `UPDATE rentals SET item_id = ?, user_id = ?, start_date = ?, end_date = ?, current_states = ?, estimated_time = ?, delivery_address = ?, created_at = ? 
	          WHERE id = ?`
	_, err := r.DB.ExecContext(ctx, query,
		rental.Item_id,
		rental.User_id,
		rental.Start_date,
		rental.End_date,
		rental.Current_state,
		rental.Esstimated_time,
		rental.Delivery_address,
		rental.Created_at,
		rental.Id,
	)
	return err
}

func (r *Rental_repo) Update_rental_start_date(ctx context.Context, id uint, startDate time.Time) error {
	query := `UPDATE rentals SET start_date = ? WHERE id = ?`
	_, err := r.DB.ExecContext(ctx, query, startDate, id)
	return err
}

func (r *Rental_repo) Update_rental_end_date(ctx context.Context, id uint, endDate time.Time) error {
	query := `UPDATE rentals SET end_date = ? WHERE id = ?`
	_, err := r.DB.ExecContext(ctx, query, endDate, id)
	return err
}

func (r *Rental_repo) Update_rental_current_state(ctx context.Context, id uint, state string) error {
	query := `UPDATE rentals SET current_states = ? WHERE id = ?`
	_, err := r.DB.ExecContext(ctx, query, state, id)
	return err
}
