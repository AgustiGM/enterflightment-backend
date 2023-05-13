package data

import (
	"awesomeProject/entities"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
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

func (repo MongoRepo) GetSongById(id int) (entities.Song, error) {
	//collection := repo.db.Collection("upvotes")
	return entities.Song{}, nil
}

func (repo MongoRepo) GetMatchById(id int) (entities.Match, error) {
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

func (repo MongoRepo) GetAllMatches() []entities.Match {
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

func (repo MongoRepo) AddMatch(match entities.Match) entities.Match {
	match.ID = rand.Int()
	match.Board = "---------"
	collection := repo.db.Collection("matches")
	insertResult, err := collection.InsertOne(context.Background(), match)
	if err != nil {
		panic(err)
	}

	_, _ = strconv.Atoi(insertResult.InsertedID.(primitive.ObjectID).Hex())

	return match
}
