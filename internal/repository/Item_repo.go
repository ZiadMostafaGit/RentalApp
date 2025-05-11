package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ZiadMostafaGit/rental-app/internal/models"
)

type Item_repo struct {
	DB *sql.DB
}
type Item_images_repo struct {
	DB *sql.DB
}

func New_item_repo(db *sql.DB) *Item_repo {
	return &Item_repo{DB: db}
}
func New_item_images_repo(db *sql.DB) *Item_images_repo {
	return &Item_images_repo{DB: db}
}

func (z *Item_images_repo) Get_item_images(ctx context.Context, id uint) ([]*models.Item_image, error) {
	query := `SELECT item_id, image_url FROM item_images WHERE item_id=?`

	res, err := z.DB.QueryContext(ctx, query, id)

	if err != nil {
		return nil, err
	}

	defer res.Close()
	images := []*models.Item_image{}

	for res.Next() {
		image := &models.Item_image{}
		if err := res.Scan(&image.Item_id, &image.Item_image); err != nil {
			return nil, err
		}
		images = append(images, image)
	}
	if len(images) == 0 {
		return nil, fmt.Errorf("no images found for this id %d", id)
	}
	return images, nil

}

func (z *Item_images_repo) Delete_item_image_by_url(ctx context.Context, itemID uint, imageURL string) error {
	query := `DELETE FROM item_images WHERE item_id = ? AND image_url = ?`

	result, err := z.DB.ExecContext(ctx, query, itemID, imageURL)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no image found with item_id %d and url %s", itemID, imageURL)
	}

	return nil
}

func (z *Item_images_repo) Delete_all_item_images(ctx context.Context, itemID uint) error {
	query := `DELETE FROM item_images WHERE item_id = ?`

	_, err := z.DB.ExecContext(ctx, query, itemID)
	return err
}

func (z *Item_images_repo) Add_item_image(ctx context.Context, image *models.Item_image) error {
	query := `INSERT INTO item_images (item_id, image_url) VALUES (?, ?)`

	_, err := z.DB.ExecContext(ctx, query, image.Item_id, image.Item_image)
	return err
}

func (i *Item_repo) Get_item(ctx context.Context, id uint) (*models.Item, error) {
	query := `SELECT id, owner_id, title, description, price, status
	          FROM items WHERE id = ?`

	item := &models.Item{}

	err := i.DB.QueryRowContext(ctx, query, id).Scan(
		&item.Id,
		&item.Owner_id,
		&item.Title,
		&item.Description,
		&item.Price,
		&item.Status,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("there is no item with id %d", id)
		}
		return nil, err
	}
	return item, nil
}

func (i *Item_repo) Insert_item(ctx context.Context, item *models.Item) error {
	query := `INSERT INTO items (owner_id, title, description, price, status)
	          VALUES (?, ?, ?, ?, ?)`

	result, err := i.DB.ExecContext(ctx, query, item.Owner_id, item.Title, item.Description, item.Price, item.Status)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	if id != int64(item.Id) {
		return fmt.Errorf("item not inserted")
	}
	return nil
}

func (i *Item_repo) Delete_item(ctx context.Context, id uint) error {
	query := `DELETE FROM items WHERE id = ?`
	result, err := i.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no item found with id %d", id)
	}
	return nil
}

func (i *Item_repo) Update_item(ctx context.Context, item *models.Item) error {
	query := `UPDATE items 
	          SET owner_id = ?, title = ?, description = ?, price = ?, status = ? 
	          WHERE id = ?`

	result, err := i.DB.ExecContext(ctx, query,
		item.Owner_id, item.Title, item.Description, item.Price, item.Status, item.Id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no item found with id %d to update", item.Id)
	}
	return nil
}

func (i *Item_repo) Update_item_title(ctx context.Context, id uint, title string) error {
	return i.updateField(ctx, "title", title, id)
}

func (i *Item_repo) Update_item_description(ctx context.Context, id uint, description string) error {
	return i.updateField(ctx, "description", description, id)
}

func (i *Item_repo) Update_item_price(ctx context.Context, id uint, price float64) error {
	return i.updateField(ctx, "price", price, id)
}

func (i *Item_repo) Update_item_status(ctx context.Context, id uint, status string) error {
	return i.updateField(ctx, "status", status, id)
}

func (i *Item_repo) Update_item_owner_id(ctx context.Context, id uint, ownerID uint) error {
	return i.updateField(ctx, "owner_id", ownerID, id)
}

func (i *Item_repo) updateField(ctx context.Context, field string, value any, id uint) error {
	query := fmt.Sprintf("UPDATE items SET %s = ? WHERE id = ?", field)
	result, err := i.DB.ExecContext(ctx, query, value, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no item found with id %d to update field %s", id, field)
	}
	return nil
}
