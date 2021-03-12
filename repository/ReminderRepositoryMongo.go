package repository

import (
	"Blibots/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type ReminderMongoRepository struct {
	mongoClient *mongo.Collection
}

func newMongo() *ReminderMongoRepository {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://discord-reminder-mongo:27017"))

	if err != nil {
		panic("Error initializing mongo " + err.Error())
	}

	database := client.Database("reminder").Collection("reminders")
	mongoDB := ReminderMongoRepository{
		mongoClient: database,
	}

	return &mongoDB
}

func (repository ReminderMongoRepository) InsertOne(r model.Reminder) (*model.Reminder, error) {
	collection := repository.mongoClient.Database().Collection("reminders")
	ctx, cancel := context. WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, r)

	if err != nil {
		fmt.Println("Error inserting document", r)
		return nil, err
	}

	return &r, nil
}

func (repository ReminderMongoRepository) InsertMany(r []model.Reminder) ([]model.Reminder, error) {
	collection := repository.mongoClient.Database().Collection("reminders")
	ctx, cancel := context. WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	infs := make([]interface{}, len(r))

	for i := 0; i < len(r);i++ {
		infs[i] = r[i]
	}

	_, err := collection.InsertMany(ctx, infs)

	if err != nil {
		fmt.Println("Error inserting document", r)
		return []model.Reminder{}, err
	}

	return r, nil
}

func (repository ReminderMongoRepository) FindWithRemindTimeLessThanNow() []model.Reminder {
	var res []model.Reminder
	collection := repository.mongoClient.Database().Collection("reminders")
	ctx, cancel := context. WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.D{{"remindtime", bson.D{{"$lte", time.Now()}}}})
	defer cur.Close(ctx)

	if err != nil {
		fmt.Println("error gan")
	}

	for cur.Next(ctx) {
		var result model.Reminder
		err := cur.Decode(&result)
		if err != nil { log.Fatal(err) }
		res = append(res, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return res

}

func (repository ReminderMongoRepository) DeleteWithRemindTimeLessThanNow() {
	collection := repository.mongoClient.Database().Collection("reminders")
	ctx, cancel := context. WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection.DeleteMany(ctx, bson.D{{"remindtime", bson.D{{"$lte", time.Now()}}}})

}

func (repository ReminderMongoRepository) Close() {
	repository.Close()
}