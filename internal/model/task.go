package model

import "time"

type Task struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	// AssigneeID  int       `json:"assignee_id,omitempty"` // Optional, for tasks assigned to other users
	// Labels      []string  `json:"labels,omitempty"` // Optional, for task categorization
	// Attachments []string  `json:"attachments,omitempty"` // Optional, for file attachments
	// Deadline    *time.Time `json:"deadline,omitempty"` // Optional, for task deadlines
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}