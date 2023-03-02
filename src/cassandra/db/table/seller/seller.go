package seller

import (
	"ServerDrivenUI/src/cassandra/db/table"
	"github.com/gocql/gocql"
	"log"
)

const (
	seller_id   = "seller_id"
	seller_name = "seller_name"
)

type Seller struct {
	Id   gocql.UUID
	Name string
}

func Insert(session *gocql.Session, sellerName string) {
	var category = Seller{}
	category.Name = sellerName

	if err := session.Query(`INSERT INTO `+table.TABLE_Seller+`(
		`+seller_id+`,
		`+seller_name+`
		) VALUES (?, ?)`,
		gocql.TimeUUID(),
		category.Name,
	).Exec(); err != nil {
		log.Fatal("Error! insert into Seller ::::    ", err)
	}
}
