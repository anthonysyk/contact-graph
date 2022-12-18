package store

import (
	"contacts/contact"
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestLoadData(t *testing.T) {
	nbOfNodes := 100000
	nbOfContactsPerNodes := 50
	filename := generateGraphName(nbOfNodes, nbOfContactsPerNodes)
	// open the store
	if _, err := os.Stat(filename); !errors.Is(err, os.ErrNotExist) {
		fmt.Println("contact already exists")
		return
	}
	store, err := Open(fmt.Sprintf("../tmp/%s", filename))
	checkErr(err)

	graph := contact.NewContactGraph()
	allPhones := contact.PopulateRandom(nbOfNodes, nbOfContactsPerNodes, graph)

	user1 := contact.User{PhoneNumber: allPhones[0]}

	// put: encodes value with gob and updates the boltdb
	err = store.Put("contact", graph)
	checkErr(err)

	err = store.Put("user1", user1)
	checkErr(err)

	// close the store
	store.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
