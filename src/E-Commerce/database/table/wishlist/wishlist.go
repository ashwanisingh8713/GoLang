package wishlist

import (
	"E-Commerce/database/table"
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"time"
)

// swagger:response Wishlist
type Wishlist struct {
	Wishlist_id gocql.UUID `json:"wishlist_id"`
	User_id     string     `json:"user_id"`
	Product_id  string     `json:"product_id"`
	Quantity    int        `json:"quantity"`
	Created_at  time.Time  `json:"Created_at"`
}

const (
	wishlist_id = "wishlist_id"
	user_id     = "user_id"
	product_id  = "product_id"
	quantity    = "quantity"
	created_at  = "created_at"
)

func Upsert(session *gocql.Session, userId string, productId string, pQuantity int) (bool, string) {
	var isSuccess bool = true
	var msg string = ""

	query := `UPDATE ` + table.TABLE_Wishlist + ` SET ` + quantity + ` = ?,  ` + created_at + ` = ? WHERE ` + product_id + ` = ?`
	values := []interface{}{pQuantity, time.Now(), productId}

	// prepare the upsert query
	/*query := "INSERT INTO mytable (id, column1, column2) VALUES (?, ?, ?) ON CONFLICT (id) DO UPDATE SET column1 = ?, column2 = ?"
	values := []interface{}{"my_id", "value1", "value2", "new_value1", "new_value2"}*/

	// prepare the upsert query
	//query := `INSERT INTO ` + table.TABLE_Wishlist + ` ( ` + wishlist_id + ` = ?, ` + user_id + ` = ?, ` + product_id + ` = ?, ` + quantity + ` = ?,  ` + created_at + ` = ? WHERE ` + product_id + ` = ?`
	//values := []interface{}{gocql.TimeUUID(), userId, productId, pQuantity, time.Now(), productId}

	err := session.Query(query, values).Exec()
	if err != nil {
		isSuccess = false
		msg = "Failed to save in wishlist."
		log.Fatal("Wishlist Upsert Error :: ", err)
	} else {
		isSuccess = false
		msg = "Product is saved in wishlist."
	}
	return isSuccess, msg
}

func Insert(session *gocql.Session, userId string, productId string, pQuantity int) (bool, string) {
	var isSuccess bool = true
	var msg string = ""

	if err := session.Query(`INSERT INTO `+table.TABLE_Wishlist+`(
		`+wishlist_id+`,
		`+user_id+`,
		`+product_id+`,
		`+quantity+`,
		`+created_at+`
		) VALUES (?, ?, ?, ?, ?)`,
		gocql.TimeUUID(),
		userId,
		productId,
		pQuantity,
		time.Now(),
	).Exec(); err != nil {
		isSuccess = false
		msg = "Failed to save in wishlist."
	} else {
		isSuccess = false
		msg = "Product is saved in wishlist."
	}

	return isSuccess, msg
}

func Read(session *gocql.Session, userId string) []Wishlist {
	var allWishes []Wishlist
	iter := session.Query(`SELECT `+
		wishlist_id+
		`,`+user_id+
		`,`+product_id+
		`,`+quantity+
		`,`+created_at+
		` FROM `+table.TABLE_Wishlist+
		` WHERE `+
		user_id+
		` = ? `, userId).Iter()
	var wishlist = Wishlist{}
	for iter.Scan(&wishlist.Wishlist_id, &wishlist.User_id, &wishlist.Product_id, &wishlist.Quantity,
		&wishlist.Created_at) {
		allWishes = append(allWishes, wishlist)
		fmt.Println("Wishlist : ", wishlist.Wishlist_id, wishlist.Product_id)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return allWishes
}

func Delete(session *gocql.Session, wishlistId string) (bool, string) {
	var isSuccess bool = true
	var msg string = ""
	if err := session.Query(`DELETE `+
		` FROM `+table.TABLE_Wishlist+
		` WHERE `+
		wishlist_id+
		` = ? `, wishlistId).Exec(); err != nil {
		fmt.Println("Table Wishlist Error, Delete() : - ", err)
		isSuccess = false
		msg = "Failed to remove from wishlist."
	} else {
		isSuccess = true
		msg = "Successfully removed from wishlist."
	}
	return isSuccess, msg
}
