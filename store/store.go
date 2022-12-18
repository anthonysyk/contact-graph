package store

import (
	"contacts/contact"
	"errors"
	"fmt"
)

var (
	ErrGenerateDataBeforeLoadTesting = errors.New("generate data before running load testing")
)

func generateGraphName(nbOfNodes, nbOfContactsPerNodes int) string {
	return fmt.Sprintf("contact-%v-%v.db", nbOfNodes, nbOfContactsPerNodes)
}

func LoadGraph(nbOfNodes, nbOfContactsPerNodes int) (contact.Graph, contact.User) {
	graph := contact.NewGraph()

	var user1 contact.User

	// fetch data from local storage boltdb
	filename := generateGraphName(nbOfNodes, nbOfContactsPerNodes)
	st, err := Open(fmt.Sprintf("./tmp/%s", filename))
	if err != nil {
		panic(ErrGenerateDataBeforeLoadTesting)
	}
	err = st.Get("contact", &graph)
	if err != nil {
		panic(ErrGenerateDataBeforeLoadTesting)
	}
	err = st.Get("user1", &user1)
	if err != nil {
		panic(ErrGenerateDataBeforeLoadTesting)
	}

	return graph, user1
}
