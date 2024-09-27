package model

import (
	"task-management/apperror"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Status      string             `bson:"status" json:"status"` // "todo", "in progress", "done"
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UserID      string             `bson:"userID" json:"userID"`
}

type TaskCreateRequest struct {
	Title       string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	Status      string `bson:"status" json:"status"` // "todo", "in progress", "done"
	UserID      string `bson:"userID" json:"-"`
}

type TaskUpdateRequest struct {
	Title       string `json:"title"`       // Title of the task
	Description string `json:"description"` // Description of the task
	Status      string `json:"status"`      // Status of the task  "todo", "in progress", "done"
}

// Validate method for TaskCreateRequest
func (t *TaskCreateRequest) Validate() error {
	return validateTask(t.Title, t.Description, t.Status)
}

// Validate method for TaskUpdateRequest
func (t *TaskUpdateRequest) Validate() error {
	return validateTask(t.Title, t.Description, t.Status)
}

// Common validation function
func validateTask(title, description, status string) error {
	if title == "" {
		return apperror.ErrBadRequest.Customize("title is required")
	}
	if description == "" {
		return apperror.ErrBadRequest.Customize("description is required")
	}
	validStatuses := []string{"todo", "in progress", "done"}
	isValidStatus := false
	for _, s := range validStatuses {
		if status == s {
			isValidStatus = true
			break
		}
	}
	if !isValidStatus {
		return apperror.ErrBadRequest.Customize("invalid status: must be one of 'todo', 'in progress', or 'done'")
	}
	return nil
}

type MarkDoneRequest struct {
	TaskIDs []string `json:"task_ids"`
}
