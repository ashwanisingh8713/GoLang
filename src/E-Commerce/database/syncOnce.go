package database

import (
	"ServerDrivenUI/src/cassandra/db"
	"fmt"
	"github.com/gocql/gocql"
	"sync"
)

var once sync.Once

var DbInstance *gocql.Session

func CreateDBSession() *gocql.Session {
	if DbInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating KEYSPACE instance now.")
				cluster := gocql.NewCluster(db.HOST_IP)
				cluster.Keyspace = db.KEYSPACE
				cluster.Consistency = gocql.Quorum
				DbInstance, _ = cluster.CreateSession()
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return DbInstance
}
