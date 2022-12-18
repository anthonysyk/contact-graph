package main

import (
	"contacts/contact"
	"contacts/store"
	"fmt"
	"time"
)

func main() {
	graph, user1 := store.LoadGraph(100, 50)

	// Test Lookup
	users, duration := WithDurationLookup(graph.Lookup, user1.PhoneNumber)
	fmt.Printf("Lookup returned %v users \n", len(users))
	fmt.Printf("Lookup took : %s \n", duration)

	// Test RLookup
	users, duration = WithDurationLookup(graph.RLookup, user1.PhoneNumber)
	fmt.Printf("RLookup returned %v users \n", len(users))
	fmt.Printf("RLookup took : %s \n", duration)

	// Test Suggest
	suggestions, duration := WithDurationSuggestion(graph.Suggest, user1)
	fmt.Printf("Suggestions:\n %v\n", suggestions)
	fmt.Printf("Suggest took : %s\n", duration)
}

func WithDurationLookup(function func(string) []contact.User, phone string) ([]contact.User, time.Duration) {
	startAt := time.Now()
	users := function(phone)
	return users, time.Since(startAt)
}

func WithDurationSuggestion(function func(contact.User) []contact.Suggestion, user contact.User) ([]contact.Suggestion, time.Duration) {
	startAt := time.Now()
	suggestions := function(user)
	return suggestions, time.Since(startAt)
}
