# contact-graph

## How to run the test
```
go mod vendor
go run .
```
Specify graph dataset before running and if dataset do not exists, please generate one.

## Stack used
- Golang (currently used langage)
- BoltDB (to load pre-generated persisted data in-memory)
    - Bolt is a pure Go key/value embedded database for Go.
    - https://github.com/boltdb/bolt

## How to generate new data
- Run test **TestLoadData** in store package
- Specify nbOfNodes
- Specify nbOfContactsPerNodes

## Dataset Generation

### 100M nodes / 50 contacts
- Duration : > 3h

### 10M / 50 contacts
- Duration : > 2h

### 1M / 50 contacts
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
