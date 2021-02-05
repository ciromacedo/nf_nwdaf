package base

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func GetConnection() *mongo.Database {
	clientOptions := options.Client().ApplyURI( GetMongoDBUri())
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	conn := client.Database(GetDBName())
	return conn
}

var CollectionNames []string
func GetCollectionsName()[]string {
	result, err := GetConnection().ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for _, coll := range result {
		CollectionNames = append(CollectionNames, coll)
	}
	return CollectionNames
}

func GetNumberOfRecordsInCollection(name string) (int64, error) {
	return GetConnection().Collection(name).EstimatedDocumentCount(context.TODO())
}