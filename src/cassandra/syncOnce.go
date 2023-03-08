package main

import (
	"cassandra/db"
	"fmt"
	"github.com/gocql/gocql"
	"sync"
)

var once sync.Once

var dbInstance *gocql.Session

func createDBSession() *gocql.Session {
	if dbInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating KEYSPACE instance now.")
				cluster := gocql.NewCluster(db.HOST_IP)
				cluster.Keyspace = db.KEYSPACE
				cluster.Consistency = gocql.Quorum
				dbInstance, _ = cluster.CreateSession()
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return dbInstance
}
