package database

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/mgo.v2"
	"wuzzapcom/Coursework/api/src/common"
)

var (
	databaseName               = "BibTex"
	textbookCollectionName     = "Textbooks"
	coursesCollectionName      = "Courses"
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

func (mongo *Mongo) InsertCourse(course common.Course) error {
	fmt.Println(course)
	collection := mongo.session.DB(databaseName).C(coursesCollectionName)

	return collection.Insert(course)
}

func (mongo *Mongo) GetAllCourses() ([]common.Course, error){
	collection := mongo.session.DB(databaseName).C(coursesCollectionName)

	var courses []common.Course
	err := collection.Find(nil).All(&courses)
	if err != nil {
		return nil, err
	}

	return courses, nil
}

//InsertTextbook ..
func (mongo *Mongo) InsertTextbook(textbook common.Item) error {

	fmt.Println(textbook)

	collection := mongo.session.DB(databaseName).C(textbookCollectionName)

	return collection.Insert(textbook)

}

//InsertTextbooks ..
func (mongo *Mongo) InsertTextbooks(textbooks common.Items) error {

	for _, textbook := range textbooks {
		err := mongo.InsertTextbook(textbook)
		if err != nil {
			return err
		}
	}

	return nil

}

//FindAllTextbooks ..
func (mongo *Mongo) FindAllTextbooks() (common.Items, error) {

	collection := mongo.session.DB(databaseName).C(textbookCollectionName)

	var items common.Items
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
