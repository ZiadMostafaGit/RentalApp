package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ZiadMostafaGit/rental-app/internal/models"
)

type Category_repo struct {
	DB *sql.DB
}

func New_category_repo(db *sql.DB) *Category_repo {

	return &Category_repo{
		DB: db,
	}

}

func (c *Category_repo) Get_categories(ctx context.Context) ([]*models.Categories, error) {
	query := `
SELECT id, name, created_at, updated_at FROM categories
`
	rows, err := c.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query categories: %w", err)
	}
	defer rows.Close()
	var categories []*models.Categories
	for rows.Next() {
		var category models.Categories
		if err := rows.Scan(&category.Id, &category.Name); err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categories = append(categories, &category)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return categories, nil
}

func (c *Category_repo) get_category_by_id(ctx context.Context, id uint) (*models.Categories, error) {
	query := `
SELECT id, name, created_at, updated_at FROM categories WHERE id = ?`
	row := c.DB.QueryRowContext(ctx, query, id)
	var category models.Categories
	if err := row.Scan(&category.Id, &category.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("category with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to scan category: %w", err)
	}
	return &category, nil
}
