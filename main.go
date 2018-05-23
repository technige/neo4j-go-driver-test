package main

import (
	"fmt"
	neo4j "github.com/neo4j/neo4j-go-driver"
)

func main() {

	// Create a new driver
	driver, err := neo4j.GraphDatabaseDriver("bolt://localhost:7687")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", driver.Target())

	// Create a session
	session, err := driver.Session(neo4j.ReadAccess)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", session.Driver().Target())

	// "Run some Cypher"
	result, err := session.Run("RETURN 1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", result)

	session.Close()

	driver.Close()

}
