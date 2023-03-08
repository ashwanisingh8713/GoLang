package address

import (
	"cassandra/db/table"
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
	AddressType  string
	AddressLine1 string
	AddressLine2 string
	IsPreferred  bool
	Zip          string
	City         string
	State        string
	Country      string
	Mobile1      string
	Mobile2      string
}

func Insert(session *gocql.Session, address Address) {
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
		`+Mobile2+`
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		gocql.TimeUUID(),
		address.UserId,
		address.AddressType,
		address.AddressLine1,
		address.AddressLine2,
		address.IsPreferred,
		address.Zip,
		address.City,
		address.State,
		address.Country,
		address.Mobile1,
		address.Mobile2,
	).Exec(); err != nil {
		log.Fatal(err)
	}
}

func Read(session *gocql.Session, userId string) []interface{} {
	var addressArray []interface{}
	var address = Address{}

	iter := session.Query(`SELECT ` +
		AddressId +
		`,` + UserId +
		`,` + AddressType +
		`,` + AddressLine1 +
		`,` + AddressLine2 +
		`,` + IsPreferred +
		`,` + Zip +
		`,` + City +
		`,` + State +
		`,` + Country +
		`,` + Mobile1 +
		`,` + Mobile2 +
		` FROM ` + table.TABLE_Address).Iter()

	/*iter := session.Query(` SELECT
	    			* FROM `+
			table.TABLE_Address+
			` WHERE `+
			UserId+` = ?`,
			userId).Iter()*/

	for iter.Scan(&address.AddressId, &address.UserId, &address.AddressType, &address.AddressLine1, &address.AddressLine2,
		&address.IsPreferred, &address.Zip, &address.City, &address.State, &address.Country,
		&address.Mobile1, &address.Mobile2) {
		fmt.Println("Address Success :: ", address.AddressId, address.Zip)
		addressArray = append(addressArray, address)
	}
	if err := iter.Close(); err != nil {
		log.Fatal("Address Error :: ", err)
	}
	return addressArray
}
