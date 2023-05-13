package data

import (
	"awesomeProject/entities"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
)

type MatchRepo interface {
	GetMatchById(id int)
	GetAllMatches()
	AddMatch(game entities.Match)
}

type InMemoryMatchRepo struct {
	Counter  int
	GameList []entities.Match
}

var uri string = "mongodb://admin:1234@localhost:27017"

func (repo MongoMatchRepo) GetMatchById(id int) (entities.Match, error) {
	collection := repo.db.Collection("matches")

	var match entities.Match
	filter := bson.M{"id": id}
	err := collection.FindOne(context.Background(), filter).Decode(&match)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return entities.Match{}, nil
		} else {
			return entities.Match{}, err
		}
	}

	return match, nil
}

func (repo MongoMatchRepo) GetAllMatches() []entities.Match {
	collection := repo.db.Collection("matches")

	var matches []entities.Match
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		// handle error
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var match entities.Match
		err := cur.Decode(&match)
		if err != nil {
			// handle error
		}
		matches = append(matches, match)
	}

	if err := cur.Err(); err != nil {
		// handle error
	}

	return matches
}

func (repo MongoMatchRepo) AddMatch(match entities.Match) entities.Match {
	match.ID = repo.Counter
	repo.Counter = repo.Counter + 1

	collection := repo.db.Collection("matches")
	insertResult, err := collection.InsertOne(context.Background(), match)
	if err != nil {
		panic(err)
	}

	match.ID, _ = strconv.Atoi(insertResult.InsertedID.(primitive.ObjectID).Hex())

	return match
}
