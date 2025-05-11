package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ZiadMostafaGit/rental-app/internal/models"
)

type Messages_repo struct {
	DB *sql.DB
}

func New_messages_repo(db *sql.DB) *Messages_repo {
	return &Messages_repo{DB: db}
}

// Get all messages
func (m *Messages_repo) Get_all_messages(ctx context.Context) ([]*models.Messages, error) {
	query := `SELECT id, conversation_id, sender_id, content, created_at FROM messages`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query all messages: %w", err)
	}
	defer rows.Close()

	var messages []*models.Messages
	for rows.Next() {
		msg := &models.Messages{}
		if err := rows.Scan(&msg.Id, &msg.Conversation_id, &msg.Sender_id, &msg.Content, &msg.Created_at); err != nil {
			return nil, fmt.Errorf("scan message: %w", err)
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

// Get messages by conversation ID
func (m *Messages_repo) Get_messages_by_conversation_id(ctx context.Context, conversationID uint) ([]*models.Messages, error) {
	query := `SELECT id, conversation_id, sender_id, content, created_at FROM messages WHERE conversation_id = ?`

	rows, err := m.DB.QueryContext(ctx, query, conversationID)
	if err != nil {
		return nil, fmt.Errorf("query messages by conversation_id: %w", err)
	}
	defer rows.Close()

	var messages []*models.Messages
	for rows.Next() {
		msg := &models.Messages{}
		if err := rows.Scan(&msg.Id, &msg.Conversation_id, &msg.Sender_id, &msg.Content, &msg.Created_at); err != nil {
			return nil, fmt.Errorf("scan message: %w", err)
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

// Insert a new message
func (m *Messages_repo) Insert_message(ctx context.Context, msg *models.Messages) error {
	query := `INSERT INTO messages (conversation_id, sender_id, content, created_at) VALUES (?, ?, ?, ?)`

	_, err := m.DB.ExecContext(ctx, query, msg.Conversation_id, msg.Sender_id, msg.Content, time.Now())
	if err != nil {
		return fmt.Errorf("insert message: %w", err)
	}
	return nil
}

// Update an existing message
func (m *Messages_repo) Update_message(ctx context.Context, msg *models.Messages) error {
	query := `UPDATE messages SET content = ? WHERE id = ?`

	res, err := m.DB.ExecContext(ctx, query, msg.Content, msg.Id)
	if err != nil {
		return fmt.Errorf("update message: %w", err)
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no message found with id %d", msg.Id)
	}
	return nil
}

// Delete a message by ID
func (m *Messages_repo) Delete_message_by_id(ctx context.Context, id uint) error {
	query := `DELETE FROM messages WHERE id = ?`

	res, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete message: %w", err)
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no message found with id %d", id)
	}
	return nil
}
