package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
)

func main() {
	cassandraSetup()
}

func cassandraSetup() {
	//cluster := gocql.NewCluster("192.168.1.1", "192.168.1.2", "192.168.1.3")
	//cluster := gocql.NewCluster("127.0.0.1", "127.0.0.1", "127.0.0.1")
	cluster := gocql.NewCluster("192.168.1.6", "192.168.1.6", "192.168.1.6")
	cluster.Keyspace = "school"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	if err := session.Query(`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`,
		"me", gocql.TimeUUID(), "hello world").Exec(); err != nil {
		log.Fatal(err)
	}

	var id gocql.UUID
	var text string

	if err := session.Query(`SELECT id, text FROM tweet WHERE timeline = ? LIMIT 1`,
		"me").Consistency(gocql.One).Scan(&id, &text); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tweet:", id, text)

	iter := session.Query(`SELECT id, text FROM tweet WHERE timeline = ?`, "me").Iter()
	for iter.Scan(&id, &text) {
		fmt.Println("Tweet:", id, text)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}
