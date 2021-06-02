package data

import (
	//"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/shashank404error/shashankMongo"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func IsUserRegistered(docID string) (bool,error) {
	var isRegistered bool

	collectionName := shashankMongo.DatabaseName.Collection("input-user")
	id, _ := primitive.ObjectIDFromHex(docID)
	filter := bson.M{"_id": id}

	type getData struct {
		Phonenumber string `json:"phonenumber"`
		UserName string `json:"username"`
	}

	var document getData

	err:= collectionName.FindOne(shashankMongo.CtxForDB, filter).Decode(&document)
	if err != nil {
		log.Error("GetInitialConversationMongo ERROR:")
		log.Error(err)
	}

	if err!=nil {
		isRegistered = false
	} else {
		if document.UserName == "" {
			isRegistered = false
		} else {
			isRegistered = true
		}
	}
	return isRegistered,err
}

func AddArchiveToMongoDocument(d *CreateArchiveRequest) error {
	collectionName := shashankMongo.DatabaseName.Collection("archive")
	_, insertErr := collectionName.InsertOne(shashankMongo.CtxForDB, d)
	if insertErr != nil {
		log.Error("AddArchiveToMongoDocument ERROR:")
		log.Error(insertErr)
	} 

	return insertErr
}

func GetArchiveFromMongoDocument(d *GetArchiveRequestPagewise) (PagewiseArchiveReponseFromDB,error) {

	var document []PagewiseArchiveReponseFromDB
	var response PagewiseArchiveReponseFromDB
	collectionName := shashankMongo.DatabaseName.Collection("archive")

	options := options.Find()
	options.SetSort(bson.D{{"rankingindex", -1}})
	options.SetSkip(d.Page-1)
	options.SetLimit(1)

	cursor, err := collectionName.Find(shashankMongo.CtxForDB, bson.M{"businessid":d.BusinessID,"userid":d.UserID},options)
	if err != nil {
		log.Error("AllAgentsByBusinessID Read ERROR : ")
		log.Error(err)
	}
	if err = cursor.All(shashankMongo.CtxForDB, &document); err != nil {
		log.Error("AllAgentsByBusinessID Write ERROR : ")
		log.Error(err)
	}

	if len(document)>0 {
		response = document[0]
	}
	return response,err
}