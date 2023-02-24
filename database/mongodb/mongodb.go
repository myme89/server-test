package mongodb

import (
	"context"
	"fmt"
	"server-test/config"
	"server-test/model"
	"time"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var clientMongo *mongo.Client

// var (
// 	collectionDB string
// 	host         string
// 	port         string
// 	dbName       string
// )

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func InitMongoDB(config *config.Config) *mongo.Client {

	var err error
	collectionDB := config.Sever.ServerMongoDB.DBcollection
	dbName := config.Sever.ServerMongoDB.DBName
	host := config.Sever.ServerMongoDB.DBHost
	port := config.Sever.ServerMongoDB.DBPort

	connStr := fmt.Sprintf("mongodb://%s:%s", host, port)
	clientOpts := options.Client().ApplyURI(connStr)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientMongo, err = mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal("cannot connect to mongo db :", err)
	}

	collectionNames, err := clientMongo.Database(dbName).ListCollectionNames(
		context.TODO(),
		bson.D{})
	if err != nil {
		log.Error(err)
	}

	if !contains(collectionNames, collectionDB) {
		err := clientMongo.Database((dbName)).CreateCollection(context.TODO(), collectionDB)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = clientMongo.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Cannot ping to mongo server :", err)
	} else {
		log.Info(" Connect MongoDB success ")
	}
	return clientMongo
}

func GetAllInfo(config *config.Config) ([]model.DataInfo, error) {

	collectionDB := config.Sever.ServerMongoDB.DBcollection
	dbName := config.Sever.ServerMongoDB.DBName
	collection := clientMongo.Database(dbName).Collection(collectionDB)
	var arr []model.DataInfo
	queryString := bson.D{}
	option := options.Find()
	var err error

	cursor, err := collection.Find(context.TODO(), queryString, option)
	if err = cursor.All(context.TODO(), &arr); err != nil {
		return arr, err
	}
	return arr, err
}

func AddInfo(config *config.Config, id int, name, fullName string) error {
	collectionDB := config.Sever.ServerMongoDB.DBcollection
	dbName := config.Sever.ServerMongoDB.DBName
	collection := clientMongo.Database(dbName).Collection(collectionDB)

	insert := bson.D{
		{Key: "name", Value: name},
		{Key: "full_name", Value: fullName},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, insert)
	if err != nil {
		return err
	}
	return nil
}
