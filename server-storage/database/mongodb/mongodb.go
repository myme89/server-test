package mongodb

import (
	"context"
	"fmt"
	"server-test/server-storage/config"
	"server-test/server-storage/model"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	// err = createCollection(dbName, "User")
	// if err != nil {
	// 	log.Fatal("Cannot create collection User DB", err)
	// }

	err = createCollection(dbName, "FileUpload")
	if err != nil {
		log.Fatal("Cannot create collection User DB", err)
	}

	// err = createCollection(dbName, "TemplateInfoPerson")
	// if err != nil {
	// 	log.Fatal("Cannot create collection Template Info Person", err)
	// }

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

func FilterInfoUser(dbName string, collectionName string, userName string) ([]model.UserInfo, error) {

	collection := clientMongo.Database(dbName).Collection(collectionName)

	var arr []model.UserInfo

	queryString := bson.D{{Key: "user_name", Value: userName}}
	option := options.Find()
	var err error

	cursor, err := collection.Find(context.TODO(), queryString, option)

	if err = cursor.All(context.Background(), &arr); err != nil {
		return arr, err
	}

	fmt.Println("nhatnt", arr)
	return arr, err
}

func FilterIdFile(config *config.Config, fileName string) (string, error) {

	collectionDB := "FileUpload"
	dbName := config.Sever.ServerMongoDB.DBName
	collection := clientMongo.Database(dbName).Collection(collectionDB)

	var arr []model.FileInfoUpload

	queryString := bson.D{{Key: "file_name", Value: fileName}}
	option := options.Find()
	var err error

	cursor, err := collection.Find(context.TODO(), queryString, option)

	if err = cursor.All(context.Background(), &arr); err != nil {
		if len(arr) > 0 {
			return arr[0].Id, err
		} else {
			return "", err
		}

	}
	return arr[0].Id, err
}

func GetHashPassword(config *config.Config, useName string) model.UserInfo {

	collectionDB := "User"
	dbName := config.Sever.ServerMongoDB.DBName

	arr, _ := FilterInfoUser(dbName, collectionDB, useName)
	var infoAcc model.UserInfo
	fmt.Println(arr)
	// var hashPassword, lastName, firstName string
	if len(arr) > 0 {
		// hashPassword = arr[0].Password
		// lastName = arr[0].LastName
		// firstName = arr[0].FistName

		infoAcc = model.UserInfo{
			Id:       arr[0].Id,
			UserName: arr[0].UserName,
			Password: arr[0].Password,
			LastName: arr[0].LastName,
			FistName: arr[0].FistName,
		}
	} else {
		infoAcc = model.UserInfo{
			Id:       "",
			UserName: "",
			Password: "",
			LastName: "",
			FistName: "",
		}
		log.Error("User don't exit in database")
	}

	return infoAcc

}

func AddInfoUploadFile(config *config.Config, fileName, idUser, typeFile, link, hashContent string, size float32) error {

	// collectionDB := config.Sever.ServerMongoDB.DBcollection
	collectionDB := "FileUpload"
	dbName := config.Sever.ServerMongoDB.DBName
	collection := clientMongo.Database(dbName).Collection(collectionDB)
	fmt.Println("nhatnt check collection ")

	now := time.Now()

	insert := bson.D{
		{Key: "file_name", Value: fileName},
		{Key: "id_user", Value: idUser},
		{Key: "type_file", Value: typeFile},
		{Key: "size", Value: size},
		{Key: "create_at", Value: now.Format("2006-01-02 15:04:05")},
		{Key: "link", Value: link},
		{Key: "check_sum", Value: hashContent},
		{Key: "status", Value: "Complete Upload"},
		{Key: "status_processing", Value: "Non-performing"},
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

func GetListFile(config *config.Config, idUser string) ([]model.FileInfoUpload, error) {

	// collectionDB := config.Sever.ServerMongoDB.DBcollection
	collectionDB := "FileUpload"
	dbName := config.Sever.ServerMongoDB.DBName
	collection := clientMongo.Database(dbName).Collection(collectionDB)
	var arr []model.FileInfoUpload
	queryString := bson.D{{Key: "id_user", Value: idUser}}
	option := options.Find()
	var err error

	cursor, err := collection.Find(context.TODO(), queryString, option)
	if err = cursor.All(context.TODO(), &arr); err != nil {
		log.Error("Get List File error err: ", err)
		return arr, err
	}

	return arr, err
}

func UpdateStatus(config *config.Config, idFile, status string) error {

	// collectionDB := config.Sever.ServerMongoDB.DBcollection
	collectionDB := "FileUpload"
	dbName := config.Sever.ServerMongoDB.DBName
	collection := clientMongo.Database(dbName).Collection(collectionDB)

	fmt.Println("status ", idFile)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, _ := primitive.ObjectIDFromHex(idFile)
	filter := bson.D{{Key: "_id", Value: objectID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "status_processing", Value: status}}}}
	// updateOptions := options.Update().SetUpsert(true)
	_, err := collection.UpdateOne(
		ctx,
		filter,
		update,
		// updateOptions,
	)
	if err != nil {
		log.Info(err)
	}

	return err
}

func GetDirFile(config *config.Config, idFile string) (string, string, string, error) {

	// collectionDB := config.Sever.ServerMongoDB.DBcollection
	collectionDB := "FileUpload"
	dbName := config.Sever.ServerMongoDB.DBName
	collection := clientMongo.Database(dbName).Collection(collectionDB)

	fmt.Println("status ", idFile)

	objectID, _ := primitive.ObjectIDFromHex(idFile)
	filter := bson.D{{Key: "_id", Value: objectID}}

	option := options.Find()
	var err error
	var arr []model.FileInfoUpload

	cursor, err := collection.Find(context.TODO(), filter, option)

	if err = cursor.All(context.Background(), &arr); err != nil {
		return arr[0].Link, arr[0].FileName, arr[0].CheckSum, err
	}

	if len(arr) > 0 {
		return arr[0].Link, arr[0].FileName, arr[0].CheckSum, err
	} else {
		return "", "", "", err
	}
}

func GetShortInfoFile(config *config.Config, idFile string) ([]model.FileInfoUpload, error) {

	// collectionDB := config.Sever.ServerMongoDB.DBcollection
	collectionDB := "FileUpload"
	dbName := config.Sever.ServerMongoDB.DBName
	collection := clientMongo.Database(dbName).Collection(collectionDB)
	var arr []model.FileInfoUpload

	objectID, _ := primitive.ObjectIDFromHex(idFile)
	queryString := bson.D{{Key: "_id", Value: objectID}}
	option := options.Find()
	var err error

	cursor, err := collection.Find(context.TODO(), queryString, option)

	if err = cursor.All(context.TODO(), &arr); err != nil {
		log.Error("Get List File error err: ", err)
		return arr, err
	}
	// fmt.Println(arr)
	return arr, err
}
