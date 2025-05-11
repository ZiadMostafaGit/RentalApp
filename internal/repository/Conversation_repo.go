package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ZiadMostafaGit/rental-app/internal/models"
)

type Conversation_repo struct {
	DB *sql.DB
}

func New_conversation_repo(db *sql.DB) *Conversation_repo {
	return &Conversation_repo{
		DB: db,
	}
}
func (c *Conversation_repo) Get_conversations(ctx context.Context) ([]*models.Conversations, error) {
	query := `
SELECT id, item_id, sender_id, recever_id FROM conversations
`
	rows, err := c.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var conversations []*models.Conversations
	for rows.Next() {
		var conversation models.Conversations
		if err := rows.Scan(&conversation.Id, &conversation.Item_id, &conversation.Sender_id, &conversation.Recever_id); err != nil {
			return nil, err
		}
		conversations = append(conversations, &conversation)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return conversations, nil

}

func (c *Conversation_repo) Get_conversations_id_by_pk(ctx context.Context, sender_id, recever_id, item_id uint) (uint, error) {
	query := `
SELECT id FROM conversations WHERE sender_id = ? AND recever_id = ? AND item_id = ?`
	row := c.DB.QueryRowContext(ctx, query, sender_id, recever_id, item_id)
	var id uint
	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("conversation with sender_id %d, recever_id %d and item_id %d not found", sender_id, recever_id, item_id)
		}
		return 0, fmt.Errorf("failed to scan conversation: %w", err)
	}
	return id, nil
}

func (c *Conversation_repo) Get_conversations_by_id(ctx context.Context, id uint) (*models.Conversations, error) {
	query := `
SELECT id, item_id, sender_id, recever_id FROM conversations WHERE id = ?`
	row := c.DB.QueryRowContext(ctx, query, id)
	var conversation models.Conversations
	if err := row.Scan(&conversation.Id, &conversation.Item_id, &conversation.Sender_id, &conversation.Recever_id); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("conversation with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to scan conversation: %w", err)
	}
	return &conversation, nil
}

func (c *Conversation_repo) make_new_conversaion(ctx context.Context, item_id, sender_id, recever_id uint) (uint, error) {
	query := `
INSERT INTO conversations (item_id, sender_id, recever_id) VALUES (?, ?, ?)`
	result, err := c.DB.ExecContext(ctx, query, item_id, sender_id, recever_id)
	if err != nil {
		return 0, fmt.Errorf("failed to insert conversation: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}
	return uint(id), nil
}

func (c *Conversation_repo) Delete_conversation(ctx context.Context, id uint) error {
	query := `
DELETE FROM conversations WHERE id = ?`
	_, err := c.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete conversation: %w", err)
	}
	return nil
}
