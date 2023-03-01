package address

import (
	"ServerDrivenUI/src/cassandra/db/table"
	"fmt"
	"github.com/gocql/gocql"
	"log"
)

const (
	AddressId    = "address_id"
	UserId       = "user_id"
	AddressType  = "address_type"
	AddressLine1 = "address_line1"
	AddressLine2 = "address_line2"
	IsPreferred  = "is_preferred"
	Zip          = "zip"
	City         = "city"
	State        = "state"
	Country      = "country"
	Mobile1      = "mobile1"
	Mobile2      = "mobile2"
)

type Address struct {
	AddressId    string
	UserId       string
	addressType  string
	addressLine1 string
	addressLine2 string
	isPreferred  bool
	zip          string
	city         string
	state        string
	country      string
	mobile1      string
	mobile2      string
}

func Insert(session gocql.Session, address Address) {
	if err := session.Query(`INSERT INTO `+table.TABLE_Address+`(
		`+AddressId+`,
		`+UserId+`,
		`+AddressType+`,
		`+AddressLine1+`,
		`+AddressLine2+`,
		`+IsPreferred+`,
		`+Zip+`,
		`+City+`,
		`+State+`,
		`+Country+`,
		`+Mobile1+`,
		`+Mobile2+`,
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		gocql.TimeUUID(),
		address.UserId,
		address.addressType,
		address.addressLine1,
		address.addressLine2,
		address.isPreferred,
		address.zip,
		address.city,
		address.state,
		address.country,
		address.mobile1,
		address.mobile2,
	).Exec(); err != nil {
		log.Fatal(err)
	}
}

func Read(session gocql.Session, userId string) {
	var address = Address{}
	iter := session.Query(` SELECT 
    			* FROM `+
		table.TABLE_Address+
		` WHERE `+
		UserId+` = ? `,
		userId).Iter()

	for iter.Scan(&address.AddressId, &address.UserId, &address.isPreferred, &address.zip, &address.addressLine1,
		&address.addressLine2, &address.addressType, &address.city, &address.country, &address.mobile1, &address.mobile2, &address.state) {
		fmt.Println("Address : ", address.AddressId, address.zip)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}
