package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ZiadMostafaGit/rental-app/internal/models"
)

type ReviewsRepo struct {
	DB *sql.DB
}

func NewReviewsRepo(db *sql.DB) *ReviewsRepo {
	return &ReviewsRepo{DB: db}
}

// Insert a new review
func (r *ReviewsRepo) InsertReview(ctx context.Context, review *models.Review) error {
	query := `
		INSERT INTO reviews (item_id, user_id, rating, comment, created_at)
		VALUES (?, ?, ?, ?, ?)
	`

	_, err := r.DB.ExecContext(
		ctx, query,
		review.ItemID,
		review.UserID,
		review.Rating,
		review.Comment,
		time.Now(),
	)

	if err != nil {
		return fmt.Errorf("failed to insert review: %w", err)
	}
	return nil
}

// Get a review by ID
func (r *ReviewsRepo) GetReviewByID(ctx context.Context, id uint) (*models.Review, error) {
	query := `
		SELECT id, item_id, user_id, rating, comment, created_at
		FROM reviews
		WHERE id = ?
	`

	review := &models.Review{}
	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&review.ID,
		&review.ItemID,
		&review.UserID,
		&review.Rating,
		&review.Comment,
		&review.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("review with ID %d not found", id)
		}
		return nil, err
	}

	return review, nil
}

// Get all reviews for a specific item
func (r *ReviewsRepo) GetReviewsByItemID(ctx context.Context, itemID uint) ([]*models.Review, error) {
	query := `
		SELECT id, item_id, user_id, rating, comment, created_at
		FROM reviews
		WHERE item_id = ?
	`

	rows, err := r.DB.QueryContext(ctx, query, itemID)
	if err != nil {
		return nil, fmt.Errorf("failed to query reviews: %w", err)
	}
	defer rows.Close()

	var reviews []*models.Review
	for rows.Next() {
		review := &models.Review{}
		if err := rows.Scan(
			&review.ID,
			&review.ItemID,
			&review.UserID,
			&review.Rating,
			&review.Comment,
			&review.CreatedAt,
		); err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}

// Update an existing review
func (r *ReviewsRepo) UpdateReview(ctx context.Context, review *models.Review) error {
	query := `
		UPDATE reviews
		SET rating = ?, comment = ?
		WHERE id = ?
	`

	result, err := r.DB.ExecContext(ctx, query, review.Rating, review.Comment, review.ID)
	if err != nil {
		return fmt.Errorf("failed to update review: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no review found with ID %d", review.ID)
	}

	return nil
}

// Delete a review by ID
func (r *ReviewsRepo) DeleteReview(ctx context.Context, id uint) error {
	query := `DELETE FROM reviews WHERE id = ?`

	result, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete review: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no review found with ID %d", id)
	}

	return nil
}
