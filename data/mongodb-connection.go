package data

import (
	"awesomeProject/entities"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepo struct {
	Counter int
	client  *mongo.Client
	db      *mongo.Database
}

func NewMongoRepo(ctx context.Context, connectionString string, dbName string) (MongoRepo, error) {
	//clientOptions := options.Client().ApplyURI(connectionString)
	//client, err := mongo.Connect(ctx, clientOptions)
	//if err != nil {
	//	return nil, err
	//}

	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		AuthSource:    dbName,
		Username:      "admin",
		Password:      "1234",
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI).SetAuth(credential)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	db := client.Database(dbName)

	return MongoRepo{
		client: client,
		db:     db,
	}, nil
}

func (repo MongoRepo) Close(ctx context.Context) error {
	return repo.client.Disconnect(ctx)
}

func (repo MongoRepo) Save(cm entities.Match) {
	collection := repo.db.Collection("matches")
	filter := bson.D{{"id", cm.ID}}
	update := bson.D{{"$set", bson.D{{"board", cm.Board}, {"user2", cm.User2}, {"turn", cm.Turn}}}}
	//update = bson.D{{"$set", bson.D{{"user2", cm.User2}}}}
	//update = bson.D{{"$set", bson.D{{"turn", cm.Turn}}}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

}
