package seller

import (
	"E-Commerce/database/table"
	"fmt"
	"github.com/gocql/gocql"
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
		fmt.Println("Error! insert into Seller ::::    ", err)
	}
}
