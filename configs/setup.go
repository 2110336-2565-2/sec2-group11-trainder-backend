package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/omise/omise-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}
func CreateGridFSBucket(client *mongo.Client) *gridfs.Bucket {

	db := client.Database(EnvMongoDBName())

	// Create a new GridFS bucket
	bucket, err := gridfs.NewBucket(
		db,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Create a new GridFS bucket")
	return bucket
}

var DB *mongo.Client = ConnectDB()
var Bucket *gridfs.Bucket = CreateGridFSBucket(DB)

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(EnvMongoDBName()).Collection(collectionName)
	return collection
}

func ConnectOmise() *omise.Client {
	client, e := omise.NewClient(EnvOmisePublicKey(), EnvOmiseSecretKey())
	if e != nil {
		log.Fatal(e)
	}
	return client
}

var OmiseClient *omise.Client = ConnectOmise()
