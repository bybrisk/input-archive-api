package data

import (
	//"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/shashank404error/shashankMongo"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo/options"
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
	result, insertErr := collectionName.InsertOne(shashankMongo.CtxForDB, d)
	if insertErr != nil {
		log.Error("AddArchiveToMongoDocument ERROR:")
		log.Error(insertErr)
	} 

	return insertErr
}