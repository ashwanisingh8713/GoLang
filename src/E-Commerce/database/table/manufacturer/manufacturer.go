package manufacturer

import (
	"E-Commerce/database/table"
	"fmt"
	"github.com/gocql/gocql"
)

const (
	manufacturer_id   string = "manufacturer_id"
	manufacturer_name string = "manufacturer_name"
)

type Manufacturer struct {
	ManufacturerId   string `json:"manufacturerId"`
	ManufacturerName string `json:"manufacturerName"`
}

func CreateManufacturer(session *gocql.Session, manufacturerName string) (bool, string) {
	var uniqueId = gocql.TimeUUID()
	if err := session.Query(`INSERT INTO `+table.TABLE_Manufacturer+`(
		`+manufacturer_id+`,
		`+manufacturer_name+`
		) VALUES (?, ?)`,
		uniqueId,
		manufacturerName,
	).Exec(); err != nil {
		fmt.Println("Error! insert into Manufacturer ::::    ", err)
		return false, "Failed to create"
	}
	return true, uniqueId.String()
}

func GetAllManufacturer(session *gocql.Session) (bool, []Manufacturer) {
	var manufacturers []Manufacturer
	var user = Manufacturer{}
	iter := session.Query(`SELECT ` +
		manufacturer_id +
		`,` + manufacturer_name +
		` FROM ` + table.TABLE_Manufacturer).Iter()

	for iter.Scan(&user.ManufacturerId, &user.ManufacturerName) {
		manufacturers = append(manufacturers, user)
	}
	if err := iter.Close(); err != nil {
		return false, manufacturers
	}
	return true, manufacturers
}

func GetManufacturerById(session *gocql.Session, manufacturerId string) (bool, Manufacturer) {
	var manufacturer = Manufacturer{}
	err := session.Query(`SELECT `+
		manufacturer_id+
		`,`+manufacturer_name+
		` FROM `+table.TABLE_Manufacturer+
		` WHERE `+
		manufacturer_id+
		` = ? `, manufacturerId).Consistency(gocql.One).Scan(&manufacturer.ManufacturerId, &manufacturer.ManufacturerName)

	if err != nil {
		fmt.Println("User Table, Get User Info failed error is ", err)
		return false, manufacturer
	}
	return true, manufacturer
}
