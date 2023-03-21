package seller

import (
	"E-Commerce/database/table"
	"fmt"
	"github.com/gocql/gocql"
	"log"
)

const (
	seller_id   = "seller_id"
	seller_name = "seller_name"
)

type Seller struct {
	SellerId   string `json:"sellerId"`
	SellerName string `json:"sellerName"`
}

func CreateSeller(session *gocql.Session, sellerName string) (bool, string) {
	var uniqueId = gocql.TimeUUID()
	if err := session.Query(`INSERT INTO `+table.TABLE_Seller+`(
		`+seller_id+`,
		`+seller_name+`
		) VALUES (?, ?)`,
		uniqueId,
		sellerName,
	).Exec(); err != nil {
		fmt.Println("Error! insert into Seller ::::    ", err)
		return false, uniqueId.String()
	}
	return true, uniqueId.String()
}

func GetAllSeller(session *gocql.Session) (bool, []Seller) {
	var status = true
	var manufacturers []Seller
	var user = Seller{}
	iter := session.Query(`SELECT ` +
		seller_id +
		`,` + seller_name +
		` FROM ` + table.TABLE_Seller).Iter()

	for iter.Scan(&user.SellerId, &user.SellerName) {
		manufacturers = append(manufacturers, user)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
		status = false
	}
	return status, manufacturers
}

func GetSellerById(session *gocql.Session, sellerId string) (bool, Seller) {
	var manufacturer = Seller{}
	err := session.Query(`SELECT `+
		seller_id+
		`,`+seller_name+
		` FROM `+table.TABLE_Seller+
		` WHERE `+
		seller_id+
		` = ? `, sellerId).Consistency(gocql.One).Scan(&manufacturer.SellerId, &manufacturer.SellerName)

	if err != nil {
		fmt.Println("User Table, Get User Info failed error is ", err)
		return false, manufacturer
	}
	return true, manufacturer
}
