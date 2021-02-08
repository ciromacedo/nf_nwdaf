package base

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func GetConnection() (*mongo.Database, *mongo.Client, context.Context) {

	client, err := mongo.NewClient(options.Client().ApplyURI(GetMongoDBUri()))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	conn := client.Database(GetDBName())
	return conn, client, ctx
}

var CollectionNames []string
func GetCollectionsName()[]string {
	db, client, ctx := GetConnection()
	result, err := db.ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for _, coll := range result {
		CollectionNames = append(CollectionNames, coll)
	}
	CloseConnection(client, ctx)
	return CollectionNames
}

func GetNumberOfRecordsInCollection(name string) (int64, error) {
	db, client, ctx := GetConnection()
	qtd, error := db.Collection(name).EstimatedDocumentCount(context.TODO())
	CloseConnection(client, ctx)
	return qtd, error
}

func CloseConnection(cli*mongo.Client, ctx context.Context){
	cli.Disconnect(ctx)
}