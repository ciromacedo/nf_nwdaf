package main

import (
	"log"
	"nf_nwdaf/base"
	"os"
)

func main() {
	log.Println("init debug")
	log.Println(os.Getenv("DB_URI"))
	log.Println(base.GetMongoDBUri())
    log.Println(base.GetNumberOfRecordsInCollection("urilist"))
/*	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DB_URI")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	result, err := client.Database(base.GetDBName()).ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for _, coll := range result {
		log.Println(coll)
	}

	defer client.Disconnect(ctx)
*/

}

