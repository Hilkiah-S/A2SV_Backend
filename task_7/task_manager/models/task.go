package models

import "time"

type Task struct {
	ID          int       `json:"id" bson:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

