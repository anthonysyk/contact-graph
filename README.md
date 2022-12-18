# contact-graph

## What is a contact graph ?

A contact graph is a type of graph data structure that represents the contacts or interactions between individuals. It is often used to model complex systems or networks, such as social networks. 

In a contact graph, each node represents an individual or entity, and an edge between two nodes indicates that there is some kind of interaction or connection between them.

For example, in a social network, a contact graph can be used to identify influential individuals or to understand how information or ideas spread through the network.

Here is an example of a contact graph representing a social network:

```
   Alice
  /      \
Bob      Eve
  \      /
   Charlie
  /      \
David    Frank
```

In this example, the nodes represent individuals (Alice, Bob, Eve, Charlie, David, and Frank) and the edges represent connections between them. For example, the edge between nodes Alice and Bob indicates that Alice and Bob are connected in the social network, while the edge between nodes Charlie and David indicates that Charlie and David are connected in the social network.

## Why load testing ?

Load testing is an important part of the development and testing process for a contact graph implementation, as it helps to ensure that the contact graph is reliable, scalable, and capable of meeting the needs of its users.

## Implementation

I used the phone number as the identifier.
```golang
type User struct {
    PhoneNumber string // +1 123 456 789 00
}
```

The graph is a dictionary :
```golang
// Graph is directed cyclic
type Graph map[string][]User
```

I implemented these operations on the graph :
```golang
type Interface interface {
	Lookup(phoneNumber string) []User
	RLookup(phoneNumber string) []User
	Suggest(phoneNumber string) []User
}
```

For the Suggest algorithm, I used a score on different conditions : 
```golang
// - contacts who have user number
// - contacts of user's contacts
// - contacts with same phone country
// - contacts with same phone area
// - contacts in common (breadth-first search - 2 levels deep)
type Suggestion struct {
    User  User
    Score int
}
```

## How to run the test
```
go mod vendor
go run .
```
Specify graph dataset before running and if dataset do not exists, please generate one.

## Stack used
- Golang (currently used language)
- BoltDB (to load pre-generated persisted data in-memory)
    - Bolt is a pure Go key/value embedded database for Go.
    - https://github.com/boltdb/bolt

## How to generate new data
- Run test **TestLoadData** in store package (it can take a long time depending on the size of the graph - nodes/edges)
- Specify nbOfNodes
- Specify nbOfContactsPerNodes

## Dataset Generation

### 100M nodes / 50 contacts
- Duration : > 3h
- took too much time

### 10M nodes / 50 contacts
- Duration : > 2h
- took too much time

### 1M nodes / 50 contacts
- Duration : 7min 51sec
- Size: 1,85 Go

```
Lookup returned 50 users 
Lookup took : 764ns 
RLookup returned 41 users 
RLookup took : 718.510447ms 
Suggestions:
 [{{+33 19 03146375} 4} {{+33 50 32320240} 4} {{+33 49 07322510} 3} {{+33 57 32738436} 3} {{+33 10 71318564} 3} {{+33 58 36321918} 3} {{+33 65 05813533} 3} {{+33 67 31601130} 3} {{+33 56 85120779} 3} {{+33 94 93537433} 3}]
Suggest took : 715.096518ms
```

### 100K / 50 contacts
- Duration : 51sec
- Size: 200 Mo

```
Lookup returned 50 users 
Lookup took : 773ns 
RLookup returned 52 users 
RLookup took : 54.548192ms 
Suggestions:
 [{{+33 24 30823692} 5} {{+33 89 69843895} 4} {{+33 79 65783203} 4} {{+33 83 43037747} 4} {{+33 90 84058818} 4} {{+33 51 08318861} 4} {{+33 21 57737915} 4} {{+33 10 39649353} 4} {{+33 54 79944275} 4} {{+33 55 27711804} 4}]
Suggest took : 61.44685ms
```
