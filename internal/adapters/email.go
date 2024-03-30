package adapters

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EmailAdapter struct {
	DB *mongo.Database
}

func NewEmailAdapter(db *mongo.Database) *EmailAdapter {
	return &EmailAdapter{
		DB: db,
	}
}
func (e *EmailAdapter) AddNotification(userId string, notificationData bson.M) error {
	collection := e.DB.Collection("notifications")
	if collection == nil {
		// Create the collection if it doesn't exist
		err := e.DB.CreateCollection(context.Background(), "notifications")
		if err != nil {
			return err
		}
	}

	// Add userId to the notification data
	notificationData["userId"] = userId
	notificationData["seen"] = false
	notificationData["timestamp"] = time.Now()

	_, err := collection.InsertOne(context.Background(), notificationData)
	return err
}
func (email *EmailAdapter) GetAllNotifications(userId string) ([]bson.M, error) {
	collection := email.DB.Collection("notifications")
	if collection == nil {
		return nil, fmt.Errorf("collection not founc")
	}
	options := options.Find().SetSort(bson.D{{"timestamp", -1}})
	filter := bson.M{"userId": userId}
	cursor, err := collection.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	_, err = collection.UpdateMany(
		context.Background(),
		bson.M{"userId": userId, "seen": false},
		bson.M{"$set": bson.M{"seen": true}},
	)
	if err != nil {
		return nil, err
	}
	var notifications []bson.M
	for cursor.Next(context.Background()) {
		var notification bson.M
		err = cursor.Decode(&notification)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}

	return notifications, nil
}
