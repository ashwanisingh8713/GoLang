package address

import (
	"E-Commerce/database/table"
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

func CreateAddress(session *gocql.Session, userId string, addressType string, addressLine1 string, addressLine2 string,
	isPreferred bool, zip string, city string, state string, country string, mobile1 string, mobile2 string) (bool, string) {
	addressId := gocql.TimeUUID()
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
		addressId,
		userId,
		addressType,
		addressLine1,
		addressLine2,
		isPreferred,
		zip,
		city,
		state,
		country,
		mobile1,
		mobile2,
	).Exec(); err != nil {
		fmt.Println("Error!, Address create :: ", err)
		return false, ""
	}
	return true, addressId.String()
}

func Read(session *gocql.Session, userId string) []Address {
	var addressArray []Address
	var address = Address{}

	iter := session.Query(`SELECT `+
		AddressId+
		`,`+UserId+
		`,`+AddressType+
		`,`+AddressLine1+
		`,`+AddressLine2+
		`,`+IsPreferred+
		`,`+Zip+
		`,`+City+
		`,`+State+
		`,`+Country+
		`,`+Mobile1+
		`,`+Mobile2+
		` FROM `+table.TABLE_Address+` WHERE user_id = ?`, userId).Iter()

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

func Delete(session *gocql.Session, userId string, addressId string) (bool, string) {
	err := session.Query(`Delete from`+
		` FROM `+table.TABLE_Address+
		` WHERE `+
		UserId+
		`, `+
		AddressId+
		` = ? `, userId, addressId).Exec()

	if err != nil {
		fmt.Println("Error! Address table delete :: ", err)
		return false, addressId
	}

	return true, addressId
}

func UpdateAddress(session *gocql.Session, userId string, addressId string, addressType string, addressLine1 string, addressLine2 string,
	isPreferred bool, zip string, city string, state string, country string, mobile1 string, mobile2 string) (bool, string) {

	err := session.Query(`Update `+table.TABLE_Address+
		` Set `+
		AddressType+` = ?, `+
		AddressLine1+` = ?, `+
		AddressLine2+` = ?, `+
		IsPreferred+` = ?, `+
		Zip+` = ?, `+
		City+` = ?, `+
		State+` = ?, `+
		Country+` = ?, `+
		Mobile1+` = ?, `+
		Mobile2+` = ? `+
		` WHERE `+
		UserId+` =? AND `+
		AddressId+` = ? `,
		addressType, addressLine1, addressLine2, isPreferred, zip, city, state, country, mobile1, mobile2,
		userId, addressId).Exec()

	if err != nil {
		fmt.Println("Error! Address table update :: ", err)
		return false, addressId
	}
	return true, addressId
}
