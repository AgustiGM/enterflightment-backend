package data

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoMatchRepo struct {
	Counter int
	client  *mongo.Client
	db      *mongo.Database
}

func NewMongoMatchRepo(ctx context.Context, connectionString string, dbName string) (MongoMatchRepo, error) {
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

	return MongoMatchRepo{
		client: client,
		db:     db,
	}, nil
}

func (repo MongoMatchRepo) Close(ctx context.Context) error {
	return repo.client.Disconnect(ctx)
}
