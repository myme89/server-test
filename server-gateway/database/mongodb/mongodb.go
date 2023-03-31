package mongodb

import (
	"context"
	"fmt"
	"server-test/server-gateway/config"
	"server-test/server-gateway/model"
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
	// collectionDB := config.Sever.ServerMongoDB.DBcollection
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

	// err = createCollection(dbName, collectionDB)
	// if err != nil {
	// 	log.Fatal("Cannot create collection DB", err)
	// }

	err = createCollection(dbName, "ServiceFuntionProcess")
	if err != nil {
		log.Fatal("Cannot create collection ServiceGateway DB", err)
	}

	err = createCollection(dbName, "ServiceUploadDB")
	if err != nil {
		log.Fatal("Cannot create collection ServiceUpload DB", err)
	}

	err = createCollection(dbName, "ServiceProcessDB")
	if err != nil {
		log.Fatal("Cannot create collection ServiceProcess DB", err)
	}

	err = clientMongo.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Cannot ping to mongo server :", err)
	}

	return clientMongo
}

func GetAllInfo(config *config.Config, collectionDB string) ([]model.TemplateInfoPerson, error) {

	// collectionDB := config.Sever.ServerMongoDB.DBcollection
	dbName := config.Sever.ServerMongoDB.DBName
	collection := clientMongo.Database(dbName).Collection(collectionDB)
	var arr []model.TemplateInfoPerson
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

func GetListServiceUpload(config *config.Config) ([]model.InfoServiceUpload, error) {

	// collectionDB := config.Sever.ServerMongoDB.DBcollection
	collectionDB := "ServiceUploadDB"
	dbName := config.Sever.ServerMongoDB.DBName
	collection := clientMongo.Database(dbName).Collection(collectionDB)
	var arr []model.InfoServiceUpload
	queryString := bson.D{}
	option := options.Find()
	var err error

	cursor, err := collection.Find(context.TODO(), queryString, option)
	if err = cursor.All(context.TODO(), &arr); err != nil {
		log.Error("Get List Service Upload error err: ", err)
		return arr, err
	}
	return arr, err
}

func GetListServiceProcess(config *config.Config, idServiceUpload string) ([]model.InfoServiceProcess, error) {

	// collectionDB := config.Sever.ServerMongoDB.DBcollection
	collectionDB := "ServiceProcessDB"
	dbName := config.Sever.ServerMongoDB.DBName
	collection := clientMongo.Database(dbName).Collection(collectionDB)
	var arr []model.InfoServiceProcess
	queryString := bson.D{{Key: "id_service_upload", Value: idServiceUpload}}
	option := options.Find()
	var err error

	cursor, err := collection.Find(context.TODO(), queryString, option)
	if err = cursor.All(context.TODO(), &arr); err != nil {
		log.Error("Get List Service Upload error err: ", err)
		return arr, err
	}
	return arr, err
}

func GetListFunctionProcess(config *config.Config, idServiceProcess string) ([]model.InfoFunctionProcess, error) {

	// collectionDB := config.Sever.ServerMongoDB.DBcollection
	collectionDB := "ServiceFuntionProcess"
	dbName := config.Sever.ServerMongoDB.DBName

	collection := clientMongo.Database(dbName).Collection(collectionDB)
	var arr []model.InfoFunctionProcess
	queryString := bson.D{{Key: "id_service_process", Value: idServiceProcess}}
	option := options.Find()
	var err error

	cursor, err := collection.Find(context.TODO(), queryString, option)
	if err = cursor.All(context.TODO(), &arr); err != nil {
		log.Error("Get List Service Upload error err: ", err)
		return arr, err
	}

	fmt.Println(arr)
	return arr, err
}
