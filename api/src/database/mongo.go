package database

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/mgo.v2"
	"wuzzapcom/Coursework/api/src/bibtex"
)

var (
	databaseName               = "BibTex"
	textbookCollectionName     = "Textbooks"
	courcesCollectionName      = "Cources"
	bibliographyCollectionName = "Bibliography"
)

//MongoConfiguration ..
type MongoConfiguration struct {
	ServerAddress string
}

func (config MongoConfiguration) String() string {
	return fmt.Sprintf("MongoConfiguration: \n\tServerAddress: %s", config.ServerAddress)
}

//Mongo ..
type Mongo struct {
	Configuration *MongoConfiguration
	session       *mgo.Session
}

//Connect ..
func (mongo *Mongo) Connect() {

	if mongo.Configuration != nil {
		log.Println("Connect to MongoDB with" + mongo.Configuration.String())
	}

	var url string
	var err error
	if mongo.Configuration != nil {
		url = mongo.Configuration.ServerAddress
	} else {
		url = "127.0.0.1"
	}

	mongo.session, err = mgo.Dial(url)
	if err != nil {
		mongo.exitWithMessage(err.Error())
	}

	log.Println("Successfull connection")
}

//InsertTextbook ..
func (mongo *Mongo) InsertTextbook(textbook bibtex.Item) error {

	collection := mongo.session.DB(databaseName).C(textbookCollectionName)

	return collection.Insert(textbook)

}

//InsertTextbooks ..
func (mongo *Mongo) InsertTextbooks(textbooks bibtex.Items) error {

	for _, textbook := range textbooks {
		err := mongo.InsertTextbook(textbook)
		if err != nil {
			return err
		}
	}

	return nil

}

//FindAllTextbooks ..
func (mongo *Mongo) FindAllTextbooks() (bibtex.Items, error) {

	collection := mongo.session.DB(databaseName).C(textbookCollectionName)

	var items bibtex.Items
	err := collection.Find(nil).All(&items)
	if err != nil {
		return nil, err
	}

	return items, nil

}

//DropTextbooks ..
func (mongo *Mongo) DropTextbooks() {

	collection := mongo.session.DB(databaseName).C(textbookCollectionName)

	collection.DropCollection()

}

func (mongo Mongo) exitWithMessage(message string) {
	if mongo.session != nil {
		mongo.session.Close()
	}
	fmt.Println(message)
	log.Println(message)
	os.Exit(-1)
}

//Disconnect ..
func (mongo Mongo) Disconnect() {
	mongo.session.Close()
}
