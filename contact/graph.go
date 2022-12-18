package contact

import (
	"errors"
	"sort"
)

// Graph is directed cyclic
type Graph map[string][]User

func NewGraph() Graph {
	return make(map[string][]User)
}

func (cg Graph) addNode(phoneNumber string) {
	cg[phoneNumber] = []User{}
}

func (cg Graph) addEdge(phoneNumber string, contact User) error {
	for _, user := range cg[phoneNumber] {
		if user.PhoneNumber == contact.PhoneNumber {
			return errors.New("node2 is already a contact of node1")
		}
	}
	cg[phoneNumber] = append(cg[phoneNumber], contact)
	return nil
}

// Lookup returns all contacts of this phone number
func (cg Graph) Lookup(phoneNumber string) []User {
	return cg[phoneNumber]
}

// RLookup returns all contacts who have this phone number in their contacts
func (cg Graph) RLookup(phoneNumber string) []User {
	var usersWithPhoneNumber []User
	for phone, contacts := range cg {
		for _, user := range contacts {
			if user.PhoneNumber == phoneNumber {
				usersWithPhoneNumber = append(usersWithPhoneNumber, User{PhoneNumber: phone})
			}
		}
	}
	return usersWithPhoneNumber
}

// Suggest returns 10 suggested contacts
// - contacts who have user's number in their contacts
// - contacts of user's contacts
// - contacts with same phone country
// - contacts with same phone area
// - contacts in common (breadth-first search - 2 levels deep)
func (cg Graph) Suggest(user User) []Suggestion {
	const (
		ScoreHasUserPhoneNumber = 3
		ScoreHasCommonContact   = 2
		ScoreIsSamePhoneArea    = 2
		ScoreIsSamePhoneCountry = 1
	)
	// todo: get contacts with same phone area and add +2 score
	// todo: get contacts with same country and add +1 score

	// use RLookup to get all contacts who have my numbers
	rlookupUsers := cg.RLookup(user.PhoneNumber)

	// get all contacts with 1 common contact with user
	var commonContacts []User
	for _, contact := range cg.Lookup(user.PhoneNumber) {
		for _, _contact := range cg.Lookup(contact.PhoneNumber) {
			commonContacts = append(commonContacts, _contact)
		}
	}

	// add scores with available data to potential suggested contacts
	scoredContactsMap := make(map[User]int)

	// add scores to all contacts
	for _, u := range rlookupUsers {
		scoredContactsMap[u] += ScoreHasUserPhoneNumber
	}
	for _, u := range commonContacts {
		scoredContactsMap[u] += ScoreHasCommonContact
	}

	suggestions := OrderSuggestions(scoredContactsMap)

	// todo: add score same country
	// todo: add score same phone area

	return suggestions[:10]
}

func OrderSuggestions(suggestions map[User]int) SuggestionList {
	p := make(SuggestionList, len(suggestions))
	idx := 0
	for k, v := range suggestions {
		p[idx] = Suggestion{k, v}
		idx++
	}

	sort.Sort(p)

	return p
}
