package database

import (
	"E-Commerce/ec_constant"
	"fmt"
	"github.com/gocql/gocql"
	"sync"
)

const (
	KEYSPACE = "ecommerce"
)

var once sync.Once

var DbInstance *gocql.Session

func CreateDBSession() *gocql.Session {
	if DbInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating KEYSPACE instance now.")
				cluster := gocql.NewCluster(ec_constant.HOST)
				cluster.Keyspace = KEYSPACE
				cluster.Consistency = gocql.Quorum
				DbInstance, _ = cluster.CreateSession()
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return DbInstance
}
