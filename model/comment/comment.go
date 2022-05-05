package comment

import (
	"context"
	"fmt"
	"stranger-album-api/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Comment struct {
	// omit Id field in json, ObjectID will be generated later
	Id         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	PhotoId    string             `bson:"photo_id" json:"photo_id"`
	Creator    string             `bson:"creator" json:"creator"`
	// json data needs to follow RFC 3999 time format, "yyyy-MM-ddThh:mm:ss+08:00"
	CreateTime time.Time          `bson:"create_time" json:"create_time"`
	Content    string             `bson:"content" json:"content"`
}

func InsertOne(client *mongo.Client, database string, comment Comment) error {
	db := client.Database(database)
	coll := db.Collection(model.CommentCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// generate a new random objectID
	comment.Id = primitive.NewObjectID()
	result, err := coll.InsertOne(ctx, comment)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Inserted transaction with Id:", result.InsertedID)
	return err
}

func FindAll(client *mongo.Client, database string, photoId string) ([]Comment, error) {
	var comments []Comment
	var err error

	db := client.Database(database)
	coll := db.Collection(model.CommentCollection)

	filter := bson.M{
		"photo_id": photoId,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := coll.Find(ctx, filter)
	if err != nil {
		fmt.Println(err)
		return comments, err
	}

	defer result.Close(ctx)
	for result.Next(ctx) {
		var c Comment
		if err = result.Decode(&c); err != nil {
			fmt.Println(err)
			return []Comment{}, err
		}
		comments = append(comments, c)
	}

	fmt.Printf("Fonud comment with Id <%v>", photoId)
	return comments, err
}
