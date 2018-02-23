package database

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

type Person struct {
	Name string
	Age  int
}

func TestConn() {

	fmt.Println("fuck")

	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	collection := session.DB("Test").C("collection")

	// err = collection.DropCollection()
	// if err != nil {
	// 	panic(err)
	// }

	err = collection.Insert(&Person{Name: "Vladimir", Age: 21})
	if err != nil {
		panic(err)
	}

	var results []Person

	err = collection.Find(nil).All(&results)
	if err != nil {
		panic(err)
	}

	for _, person := range results {
		fmt.Println(person)
	}

}
