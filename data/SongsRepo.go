package data

import (
	"awesomeProject/entities"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"sort"
)

type SongRepo interface {
	GetSongById(id string)
	GetAllSongs()
	AddUpvoteSong(song entities.Song)
	GetUpvotes()
	AddUpvote(id string)
	EmptyPlaylist()
	PrepareNextSong()
}

func GetSongById(id string) (entities.Song, error) {
	for _, s := range songs {
		if s.ID == id {
			return s, nil
		}
	}
	return entities.Song{}, nil
}

func (repo MongoRepo) GetAllSongs() ([]entities.Song, error) {
	collection := repo.db.Collection("playlist")

	var playlist []entities.Song
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		// handle error
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var song entities.Song
		err := cur.Decode(&song)
		if err != nil {
			// handle error
		}
		playlist = append(playlist, song)
	}

	var result []entities.Song
	var songs2 = songs
	for _, s := range songs2 {
		found := false
		for _, p := range playlist {
			if s.ID == p.ID {
				found = true
			}
		}
		if !found {
			result = append(result, s)
		}
	}

	return result, nil
}

func (repo MongoRepo) GetPlaylist() ([]entities.Song, error) {
	collection := repo.db.Collection("playlist")

	var playlist []entities.Song
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		// handle error
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var song entities.Song
		err := cur.Decode(&song)
		if err != nil {
			// handle error
		}
		playlist = append(playlist, song)
	}
	//-------------------playlist-------------------

	collection2 := repo.db.Collection("upvotes")
	//converter la collection en upvotes
	var upvotes []entities.Upvote
	cur2, _ := collection2.Find(context.Background(), bson.M{})
	defer cur2.Close(context.Background())
	for cur.Next(context.Background()) {
		var upvote entities.Upvote
		err := cur.Decode(&upvote)
		if err != nil {
			// handle error
		}
		upvotes = append(upvotes, upvote)
	}
	//-----------upvotes----------------

	songs2 := playlist
	sort.Slice(songs2, func(i, j int) bool {
		songID1 := songs2[i].ID
		songID2 := songs2[j].ID

		// Finding upvote objects for the respective songs
		upvote1 := findUpvote(songID1, upvotes)
		upvote2 := findUpvote(songID2, upvotes)

		// Comparing upvotes
		return upvote1.Upvotes > upvote2.Upvotes
	})

	solution := []entities.Song{}
	// Print sorted songs
	for _, song := range playlist {
		solution = append(solution, song)
	}

	if err := cur.Err(); err != nil {
		// handle error
	}
	return songs2, err
}
func findUpvote(songID string, upvotes []entities.Upvote) entities.Upvote {
	for _, upvote := range upvotes {
		if upvote.SongID == songID {
			return upvote
		}
	}
	return entities.Upvote{}
}

func (repo MongoRepo) AddSongToPlaylist(id string) ([]entities.Song, error) {

	collection := repo.db.Collection("playlist")

	//busquem la newSong a la llista de songs
	var newSong entities.Song
	for _, s := range songs {
		if s.ID == id {
			newSong = s
		}
	}

	//converter la collection en playlist
	var playlist []entities.Song
	cur, err := collection.Find(context.Background(), bson.M{})

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var song entities.Song
		err := cur.Decode(&song)
		if err != nil {
			// handle error
		}
		playlist = append(playlist, song)
	}

	//si la canço nova no existeix, no fem res
	if newSong.ID == "" {
		return playlist, nil
	}

	insertResult, err := collection.InsertOne(context.Background(), newSong)

	if err != nil {
		fmt.Println(insertResult)
		panic(err)
	}

	return append(playlist, newSong), err
}

var songs = []entities.Song{
	{
		ID:     "1",
		Title:  "Call me",
		Artist: "Blondie",
		Length: 10.00,
		Album:  "Call me",
		Year:   "1980",
	},
	{
		ID:     "2",
		Title:  "Bette Davis Eyes",
		Artist: "Kim Carnes",
		Length: 10.00,
		Album:  "Mistaken Identity",
		Year:   "1981",
	},
	{
		ID:     "3",
		Title:  "Physical",
		Artist: "Olive Newton-John",
		Length: 10.00,
		Album:  "Physical",
		Year:   "1982",
	},
	{
		ID:     "4",
		Title:  "Every Breath You Take",
		Artist: "The Police",
		Length: 10.00,
		Album:  "Synchronicity",
		Year:   "1983",
	},
	{
		ID:     "5",
		Title:  "When Doves Cry",
		Artist: "Prince",
		Length: 10.00,
		Album:  "Purple Rain",
		Year:   "1984",
	},
	{
		ID:     "6",
		Title:  "Careless Whisper",
		Artist: "George Michael",
		Length: 10.00,
		Album:  "Make It Big",
		Year:   "1985",
	},
	{
		ID:     "7",
		Title:  "That's What Friends Are For",
		Artist: "Dionne and Friends",
		Length: 10.00,
		Album:  "Friends",
		Year:   "1986",
	},
	{
		ID:     "8",
		Title:  "Walk Like an Egyptian",
		Artist: "The Bangles",
		Length: 10.00,
		Album:  "Different Light",
		Year:   "1987",
	},
	{
		ID:     "9",
		Title:  "Faith",
		Artist: "George Michael",
		Length: 10.00,
		Album:  "Faith",
		Year:   "1988",
	},
	{
		ID:     "10",
		Title:  "Roll With It",
		Artist: "Steve Winwood",
		Length: 10.00,
		Album:  "Roll With It",
		Year:   "1989",
	},
	{
		ID:     "11",
		Title:  "Another Day in Paradise",
		Artist: "Phil Collins",
		Length: 10.00,
		Album:  "But Seriously",
		Year:   "1990",
	},
	{
		ID:     "12",
		Title:  "Unbelievable",
		Artist: "EMF",
		Length: 10.00,
		Album:  "Schubert Dip",
		Year:   "1991",
	},
	{
		ID:     "13",
		Title:  "End of the Road",
		Artist: "Boyz II Men",
		Length: 10.00,
		Album:  "Boomerang",
		Year:   "1992",
	},
	{
		ID:     "14",
		Title:  "I Will Always Love You",
		Artist: "Whitney Houston",
		Length: 10.00,
		Album:  "The Bodyguard",
		Year:   "1993",
	},
	{
		ID:     "15",
		Title:  "The Sign",
		Artist: "Ace of Base",
		Length: 10.00,
		Album:  "Happy Nation",
		Year:   "1994",
	},
	{
		ID:     "16",
		Title:  "Gangsta's Paradise",
		Artist: "Coolio",
		Length: 10.00,
		Album:  "Dangerous Minds",
		Year:   "1995",
	},
	{
		ID:     "17",
		Title:  "Macarena",
		Artist: "Los del Rio",
		Length: 10.00,
		Album:  "A mi me gusta",
		Year:   "1996",
	},
	{
		ID:     "18",
		Title:  "Something About the Way You Look Tonight",
		Artist: "Elton John",
		Length: 10.00,
		Album:  "The Big Picture",
		Year:   "1997",
	},
	{
		ID:     "19",
		Title:  "Too Close",
		Artist: "Next",
		Length: 10.00,
		Album:  "Rated Next",
		Year:   "1998",
	},
	{
		ID:     "20",
		Title:  "Believe",
		Artist: "Cher",
		Length: 10.00,
		Album:  "Believe",
		Year:   "1999",
	},
	{
		ID:     "21",
		Title:  "Breathe",
		Artist: "Faith Hill",
		Length: 10.00,
		Album:  "Breathe",
		Year:   "2000",
	},
	{
		ID:     "22",
		Title:  "Hanging by a Moment",
		Artist: "Lifehouse",
		Length: 10.00,
		Album:  "No Name Face",
		Year:   "2001",
	},
}

func (repo MongoRepo) GetUpvotes() ([]entities.Upvote, error) {
	collection := repo.db.Collection("upvotes")

	//converter la collection en playlist
	var upvotes []entities.Upvote
	cur, err := collection.Find(context.Background(), bson.M{})

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var upvote entities.Upvote
		err := cur.Decode(&upvote)
		if err != nil {
			// handle error
		}
		upvotes = append(upvotes, upvote)
	}

	return upvotes, err

}

func (repo MongoRepo) AddUpvote(id string) error {

	collection := repo.db.Collection("upvotes")

	//converter la collection en upvotes
	var upvotes []entities.Upvote
	cur, err := collection.Find(context.Background(), bson.M{})

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var upvote entities.Upvote
		err := cur.Decode(&upvote)
		if err != nil {
			// handle error
		}
		upvotes = append(upvotes, upvote)
	}

	//si la canço nova no existeix, no fem res
	var newUpvote entities.Upvote
	var found = false
	for _, u := range upvotes {
		if u.SongID == id {
			found = true
			newUpvote = u
			newUpvote.Upvotes++
		}
	}
	if !found {
		newUpvote = entities.Upvote{ID: id, SongID: id, Upvotes: 1}
		insertResult, err := collection.InsertOne(context.Background(), newUpvote)
		fmt.Println(insertResult, err)
		return nil
	}

	filter := bson.M{"_songid": id}
	// Create the update to be applied
	update := bson.M{"$set": bson.M{"upvotes": newUpvote.Upvotes}}
	// Perform the update operation
	updateResult, err := collection.UpdateOne(context.Background(), filter, update)
	fmt.Println(updateResult)

	return err
}

func (repo MongoRepo) EmptyPlaylist() error {
	collection := repo.db.Collection("playlist")
	_, err := collection.DeleteMany(context.Background(), bson.M{})
	return err
}

func (repo MongoRepo) PrepareNextSong() (error, error) {
	collection := repo.db.Collection("playlist")

	var playlist []entities.Song
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		// handle error
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var song entities.Song
		err := cur.Decode(&song)
		if err != nil {
			// handle error
		}
		playlist = append(playlist, song)
	}
	if len(playlist) > 3 {
		collection.DeleteOne(context.Background(), bson.M{"_songid": playlist[0].ID})
	}

	collection2 := repo.db.Collection("upvotes")
	_, err3 := collection2.DeleteMany(context.Background(), bson.M{})

	return nil, err3
}
