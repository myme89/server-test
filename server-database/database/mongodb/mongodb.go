package mongodb

import (
	"context"
	"fmt"
	"server-test/server-database/config"
	"server-test/server-database/model"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var clientMongo *mongo.Client

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func createCollection(dbName string, collectionDB string) error {
	collectionNames, err := clientMongo.Database(dbName).ListCollectionNames(
		context.TODO(),
		bson.D{})
	if err != nil {
		log.Error("createCollection erro err: ", err)
	}

	if !contains(collectionNames, collectionDB) {
		err := clientMongo.Database((dbName)).CreateCollection(context.TODO(), collectionDB)
		if err != nil {
			log.Fatal("createCollection erro err: ", err)
		}
	}
	return err
}

func InitMongoDB(config *config.Config) *mongo.Client {

	var err error
	collectionDB := config.Sever.ServerMongoDB.DBcollection
	dbName := config.Sever.ServerMongoDB.DBName
	host := config.Sever.ServerMongoDB.DBHost
	port := config.Sever.ServerMongoDB.DBPort

	// username := config.Sever.ServerMongoDB.DBUserName
	// password := config.Sever.ServerMongoDB.DBPassword

	// credential := options.Credential{
	// 	Username: username,
	// 	Password: password,
	// }

	connStr := fmt.Sprintf("mongodb://%s:%s", host, port)
	// clientOpts := options.Client().ApplyURI(connStr).SetAuth(credential)
	clientOpts := options.Client().ApplyURI(connStr)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientMongo, err = mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal("cannot connect to mongo db :", err)
	}

	err = createCollection(dbName, collectionDB)
	if err != nil {
		log.Fatal("Cannot create collection DB", err)
	}

	err = createCollection(dbName, "User")
	if err != nil {
		log.Fatal("Cannot create collection User DB", err)
	}

	err = clientMongo.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Cannot ping to mongo server :", err)
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
		log.Error("GetAllInfo error err: ", err)
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
		log.Error("AddInfo err = ", err)
		return err
	}

	return nil
}

func AddManyInfo(config *config.Config, info []model.DataPost) error {

	collectionDB := config.Sever.ServerMongoDB.DBcollection
	dbName := config.Sever.ServerMongoDB.DBName
	collection := clientMongo.Database(dbName).Collection(collectionDB)

	infos := make([]interface{}, len(info))
	for i, s := range info {
		infos[i] = s
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertMany(ctx, infos)
	if err != nil {
		log.Error("AddManyInfo error : ", err)
		return err
	}

	return nil
}

func AddManyInfoNotModel(config *config.Config, info []interface{}, nameCollection string) error {
	// collectionDB := config.Sever.ServerMongoDB.DBcollection
	dbName := config.Sever.ServerMongoDB.DBName

	err := createCollection(dbName, nameCollection)
	if err != nil {
		log.Fatal("Cannot create collection DB", err)
	}

	collection := clientMongo.Database(dbName).Collection(nameCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = collection.InsertMany(ctx, info)
	if err != nil {
		log.Error("AddManyInfoNotModel error: ", err)
		return err
	}
	return nil
}

// func AddUserInfoSignUp(userName, password, lastName, firstName, dbName, collectionDB string) error {

// 	// collectionDB := config.Sever.ServerMongoDB.DBcollection
// 	// dbName := config.Sever.ServerMongoDB.DBName
// 	collection := clientMongo.Database(dbName).Collection(collectionDB)
// 	fmt.Println("nhatnt check collection ", collection)

// 	// insert := bson.D{
// 	// 	{Key: "user_name", Value: userName},
// 	// 	{Key: "full_name", Value: password},
// 	// 	{Key: "last_name", Value: lastName},
// 	// 	{Key: "first_name", Value: firstName},
// 	// 	{Key: "full_name", Value: firstName + " " + lastName},
// 	// }

// 	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	// defer cancel()
// 	// _, err := collection.InsertOne(ctx, insert)
// 	// if err != nil {
// 	// 	log.Error("AddInfo err = ", err)
// 	// 	return err
// 	// }

// 	return nil
// }

func AddUserInfoSignUp(config *config.Config, userName, password, lastName, firstName string) error {

	// collectionDB := config.Sever.ServerMongoDB.DBcollection
	collectionDB := "User"
	dbName := config.Sever.ServerMongoDB.DBName
	collection := clientMongo.Database(dbName).Collection(collectionDB)
	fmt.Println("nhatnt check collection ", collection)

	insert := bson.D{
		{Key: "user_name", Value: userName},
		{Key: "password", Value: password},
		{Key: "last_name", Value: lastName},
		{Key: "first_name", Value: firstName},
		{Key: "full_name", Value: firstName + " " + lastName},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, insert)
	if err != nil {
		log.Error("AddInfo err = ", err)
		return err
	}

	return nil
}
