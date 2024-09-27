package task

import (
	"context"
	"task-management/apperror"
	"task-management/consts"
	"task-management/model"
	"task-management/mongodatabase"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func taskCreate(db *mongodatabase.DBConfig, mongoClient *mongo.Client, req model.TaskCreateRequest) (string, error) {

	taskCollection := db.NewCollection(mongoClient, consts.TaskCollection)

	taskdata := model.Task{
		ID:          primitive.NewObjectID(),
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		CreatedAt:   time.Now(),
		UserID:      req.UserID,
	}

	_, err := taskCollection.InsertOne(context.TODO(), taskdata)
	if err != nil {
		return "", apperror.ErrServer.Customize(err.Error()).LogWithLocation()
	}

	return taskdata.ID.Hex(), nil
}

func taskRead(db *mongodatabase.DBConfig, mongoClient *mongo.Client, taskID string, userID string) (*model.Task, error) {
	taskCollection := db.NewCollection(mongoClient, consts.TaskCollection)

	objID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return nil, apperror.ErrUnauthorized.Customize("Invalid Task ID").LogWithLocation()
	}

	filter := bson.M{"_id": objID, "userID": userID}
	var task model.Task
	err = taskCollection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, apperror.ErrNotFound.Customize("Task not found").LogWithLocation()
		}
		return nil, apperror.ErrServer.Customize(err.Error()).LogWithLocation()
	}

	return &task, nil
}

func taskList(db *mongodatabase.DBConfig, mongoClient *mongo.Client, userID string, limit, offset int64, statusFilter *string, sortField *string, sortOrder *int) ([]model.Task, error) {
	taskCollection := db.NewCollection(mongoClient, consts.TaskCollection)

	filter := bson.M{"userID": userID}
	if statusFilter != nil {
		filter["status"] = *statusFilter
	}

	sort := bson.M{}
	if sortField != nil {
		order := 1 // Default to ascending
		if sortOrder != nil && *sortOrder == -1 {
			order = -1 // Descending order
		}
		sort[*sortField] = order
	}

	// Use FindOptions to set limit, offset, and sorting
	findOptions := options.Find().
		SetLimit(limit).
		SetSkip(offset).
		SetSort(sort)

	cursor, err := taskCollection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, apperror.ErrNotFound.Customize(err.Error()).LogWithLocation()
	}
	defer cursor.Close(context.TODO())

	var tasks []model.Task
	if err := cursor.All(context.TODO(), &tasks); err != nil {
		return nil, apperror.ErrServer.Customize(err.Error()).LogWithLocation()
	}

	return tasks, nil
}

func taskUpdate(db *mongodatabase.DBConfig, mongoClient *mongo.Client, taskID string, userID string, req model.TaskUpdateRequest) error {
	taskCollection := db.NewCollection(mongoClient, consts.TaskCollection)

	objID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return apperror.ErrBadRequest.Customize("Invalid Task ID").LogWithLocation()
	}

	filter := bson.M{"_id": objID, "userID": userID}
	update := bson.M{
		"$set": bson.M{
			"title":       req.Title,
			"description": req.Description,
			"status":      req.Status,
		},
	}

	_, err = taskCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return apperror.ErrServer.Customize(err.Error()).LogWithLocation()
	}

	return nil
}

func taskDelete(db *mongodatabase.DBConfig, mongoClient *mongo.Client, taskID string, userID string) error {
	taskCollection := db.NewCollection(mongoClient, consts.TaskCollection)

	objID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return apperror.ErrBadRequest.Customize("Invalid Task ID").LogWithLocation()
	}

	filter := bson.M{"_id": objID, "userID": userID}
	_, err = taskCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return apperror.ErrServer.Customize(err.Error()).LogWithLocation()
	}

	return nil
}

func markTasksAsDone(db *mongodatabase.DBConfig, mongoClient *mongo.Client, taskIDs []string, userID string) error {
	taskCollection := db.NewCollection(mongoClient, consts.TaskCollection)

	// Create a filter to match the tasks with the provided IDs
	filter := bson.M{
		"_id":    bson.M{"$in": taskIDs},
		"userID": userID,
	}

	// Update the tasks to set the status to "done"
	update := bson.M{"$set": bson.M{"status": "done"}}

	// Perform the update operation
	_, err := taskCollection.UpdateMany(context.Background(), filter, update)
	if err != nil {
		return apperror.ErrServer.Customize(err.Error()).LogWithLocation()
	}

	return err
}
