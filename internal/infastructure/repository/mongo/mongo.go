package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"todolist/internal/config"
	tasksentity "todolist/internal/entity/tasks"
)

type MongoDB struct {
	mongoClient *mongo.Client
	db          *mongo.Database
	collection  *mongo.Collection
}

func NewMongoDB(cfg *config.Config) (*MongoDB, error) {
	uri := "mongodb://" + cfg.Mongo.Host + cfg.Mongo.Port
	log.Printf("%s", uri)
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	return &MongoDB{
		mongoClient: client,
		db:          client.Database(cfg.Mongo.DatabaseName),
		collection:  client.Database(cfg.Mongo.DatabaseName).Collection(cfg.Mongo.CollectionName),
	}, nil
}

func (m *MongoDB) SaveDetails(ctx context.Context, req *tasksentity.MongoTaskDetails) error {
	_, err := m.collection.InsertOne(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
func (m *MongoDB) GetDetails(ctx context.Context, taskID string) (*tasksentity.MongoTaskDetails, error) {
	filter := bson.M{"task_id": taskID}

	details := &tasksentity.MongoTaskDetails{}
	err := m.collection.FindOne(ctx, filter).Decode(details)
	if err != nil {
		return nil, err
	}

	return details, nil
}

func (m *MongoDB) UpdateDetails(ctx context.Context, req *tasksentity.MongoTaskDetails) error {
	filter := bson.M{"task_id": req.TaskId}
	update := bson.M{}
	if !req.Important {
		update["important"] = req.Important
	}
	if req.Condition != "" {
		update["condition"] = req.Condition
	}
	if req.Description != "" {
		update["description"] = req.Description
	}
	updateBson := bson.M{"$set": update}
	_, err := m.collection.UpdateOne(ctx, filter, updateBson)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) DeleteDetails(ctx context.Context, taskID string) error {
	filter := bson.M{"task_id": taskID}
	_, err := m.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
