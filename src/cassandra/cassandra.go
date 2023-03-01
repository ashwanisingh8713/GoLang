package main

import (
	"ServerDrivenUI/src/cassandra/db"
	"ServerDrivenUI/src/cassandra/db/table"
	"ServerDrivenUI/src/cassandra/db/table/address"
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"sync"
)

func main() {
	cassandraSetup()
}

var once sync.Once

var dbInstance *gocql.Session

func getInstance() *gocql.Session {
	if dbInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating KEYSPACE instance now.")
				cluster := gocql.NewCluster(db.HOST_IP)
				cluster.Keyspace = db.KEYSPACE
				cluster.Consistency = gocql.Quorum
				session, _ := cluster.CreateSession()
				defer session.Close()
			})
	} else {
		fmt.Println("KEYSPACE instance already created.")
	}

	return dbInstance
}

func cassandraSetup() {
	if err := dbInstance.Query(`INSERT INTO `+table.TABLE_Address+`(
		`+address.AddressId+`,
		`+address.UserId+`,
		`+address.IsPreferred+`,
		`+address.Zip+`,
		`+address.AddressLine1+`,
		`+address.AddressType+`,
		`+address.City+`
		) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		gocql.TimeUUID(), "3343", true, "560078", "Shamanna Building", "Home", "Bangalore").Exec(); err != nil {
		log.Fatal(err)
	}

	var id gocql.UUID
	var text string

	if err := dbInstance.Query(`SELECT id, text FROM tweet WHERE timeline = ? LIMIT 1`,
		"me").Consistency(gocql.One).Scan(&id, &text); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tweet:", id, text)

	iter := dbInstance.Query(`SELECT id, text FROM tweet WHERE timeline = ?`, "me").Iter()
	for iter.Scan(&id, &text) {
		fmt.Println("Tweet:", id, text)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}
