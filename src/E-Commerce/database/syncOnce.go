package database

import (
	"fmt"
	"github.com/gocql/gocql"
	"sync"
)

const (
	HOST_IP  = "192.168.13.77"
	KEYSPACE = "ecommerce"
)

var once sync.Once

var DbInstance *gocql.Session

func CreateDBSession() *gocql.Session {
	if DbInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating KEYSPACE instance now.")
				cluster := gocql.NewCluster(HOST_IP)
				cluster.Keyspace = KEYSPACE
				cluster.Consistency = gocql.Quorum
				DbInstance, _ = cluster.CreateSession()
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return DbInstance
}
