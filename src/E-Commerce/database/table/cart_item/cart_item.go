package cart_item

import (
	"E-Commerce/database/table"
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"time"
)

const (
	cart_item_id = "cart_item_id"
	user_id      = "user_id"
	product_id   = "product_id"
	quantity     = "quantity"
	created_at   = "created_at"
	modified_at  = "modified_at"
	create_date  = "create_date"
)

type CartItem struct {
	CartItemId gocql.UUID `json:"cart_item_id"`
	UserId     string     `json:"user_id"`
	ProductId  string     `json:"product_id"`
	Quantity   int        `json:"quantity"`
	CreatedAt  time.Time  `json:"created_at"`
	ModifiedAt time.Time  `json:"modified_at"`
}

func Insert(session *gocql.Session, userId string, productId string, productQuantity int) {
	var user = CartItem{}
	user.UserId = userId
	user.ProductId = productId
	user.Quantity = productQuantity
	user.CreatedAt = time.Now()
	user.ModifiedAt = time.Now()

	if err := session.Query(`INSERT INTO `+table.TABLE_Cart_item+`(
		`+cart_item_id+`,
		`+user_id+`,
		`+product_id+`,
		`+quantity+`,
		`+created_at+`,
		`+modified_at+`
		) VALUES (?, ?, ?, ?, ?, ?)`,
		gocql.TimeUUID(),
		user.UserId,
		user.ProductId,
		user.Quantity,
		user.CreatedAt,
		user.ModifiedAt,
	).Exec(); err != nil {
		log.Fatal("Error! insert into CartItem ::::    ", err)
	}
}

func Read(session *gocql.Session, userId string) []interface{} {
	var cartItemsArray []interface{}
	var cartItem = CartItem{}

	iter := session.Query(`SELECT `+
		cart_item_id+
		`,`+user_id+
		`,`+product_id+
		`,`+quantity+
		`,`+created_at+
		`,`+modified_at+
		` FROM `+table.TABLE_Cart_item+` WHERE user_id = ? `, userId).Iter()

	/*iter := session.Query(` SELECT
	    			* FROM `+
			table.TABLE_Address+
			` WHERE `+
			UserId+` = ?`,
			userId).Iter()*/

	for iter.Scan(&cartItem.CartItemId, &cartItem.UserId,
		&cartItem.ProductId, &cartItem.Quantity,
		&cartItem.CreatedAt,
		&cartItem.ModifiedAt) {
		fmt.Println("Cart Item Success :: ", cartItem.CartItemId, cartItem.ProductId)
		cartItemsArray = append(cartItemsArray, cartItem)
	}
	if err := iter.Close(); err != nil {
		log.Fatal("Address Error :: ", err)
	}
	return cartItemsArray
}