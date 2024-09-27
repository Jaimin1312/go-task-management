package task

import (
	"task-management/model"
	"task-management/mongodatabase"

	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	TaskCreate(model.TaskCreateRequest) (string, error)
	TaskRead(taskID, userID string) (*model.Task, error)
	TaskList(userID string, limit, offset int64, statusFilter *string, sortField *string, sortOrder *int) ([]model.Task, error)
	TaskUpdate(taskID string, req model.TaskUpdateRequest, userID string) error
	TaskDelete(taskID, userID string) error
	MarkTasksAsDone(taskIDs []string, userID string) error // Add this line
}

type service struct {
	mongoDBClient *mongo.Client
	mongodb       *mongodatabase.DBConfig
}

func NewService(repo *model.Repos) Service {
	return &service{
		mongoDBClient: repo.MongoDBClient,
		mongodb:       repo.MongoDB,
	}
}

func (s *service) TaskCreate(task model.TaskCreateRequest) (string, error) {
	return taskCreate(s.mongodb, s.mongoDBClient, task)
}

func (s *service) TaskRead(taskID, userID string) (*model.Task, error) {
	return taskRead(s.mongodb, s.mongoDBClient, taskID, userID)
}

func (s *service) TaskList(userID string, limit, offset int64, statusFilter *string, sortField *string, sortOrder *int) ([]model.Task, error) {
	return taskList(s.mongodb, s.mongoDBClient, userID, limit, offset, statusFilter, sortField, sortOrder)
}

func (s *service) TaskUpdate(taskID string, req model.TaskUpdateRequest, userID string) error {
	return taskUpdate(s.mongodb, s.mongoDBClient, taskID, userID, req)
}

func (s *service) TaskDelete(taskID, userID string) error {
	return taskDelete(s.mongodb, s.mongoDBClient, taskID, userID)
}

func (s *service) MarkTasksAsDone(taskIDs []string, userID string) error {
	return markTasksAsDone(s.mongodb, s.mongoDBClient, taskIDs, userID)
}
