package contact

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func PopulateRandom(nbNodes, nbContactsPerUser int, graph Graph) []string {
	var allPhoneNumberList []string
	// generate users
	for len(graph) < nbNodes {
		fmt.Println("add node", len(graph)+1)
		phoneNumber := generateInternationalPhoneNumber()
		graph.addNode(phoneNumber)
		allPhoneNumberList = append(allPhoneNumberList, phoneNumber)
	}

	// generate contacts per user
	counter := 0
	for phone := range graph {
		counter++
		fmt.Println("add contacts to node", counter)
		for len(graph[phone]) < nbContactsPerUser {
			idx := getRandomIndex(0, len(allPhoneNumberList))
			graph.addEdge(phone, User{PhoneNumber: allPhoneNumberList[idx]})
		}
	}

	return allPhoneNumberList
}

func generateInternationalPhoneNumber() string {
	phoneNumber := gofakeit.Phone()
	number := phoneNumber[len(phoneNumber)-8:]
	area := phoneNumber[0 : len(phoneNumber)-8]
	return fmt.Sprintf("+33 %s %s", area, number)
}

func getRandomIndex(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
