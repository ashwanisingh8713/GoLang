package cart_item

import (
	"E-Commerce/database/table"
	productTable "E-Commerce/database/table/products"
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
	CartItemId gocql.UUID           `json:"cart_item_id"`
	UserId     string               `json:"user_id"`
	ProductId  string               `json:"product_id"`
	Quantity   int                  `json:"quantity"`
	CreatedAt  time.Time            `json:"created_at"`
	ModifiedAt time.Time            `json:"modified_at"`
	Product    productTable.Product `json:"product"`
}

func Insert(session *gocql.Session, userId string, productId string, productQuantity int) bool {
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
		fmt.Println("Error! insert into CartItem ::::    ", err)
		return false
	}
	return true
}

func Read(session *gocql.Session, userId string) []CartItem {
	var cartItemsArray []CartItem
	var productIds []string
	var cartItem = CartItem{}

	iter := session.Query(`SELECT `+
		cart_item_id+
		`,`+user_id+
		`,`+product_id+
		`,`+quantity+
		`,`+created_at+
		`,`+modified_at+
		` FROM `+table.TABLE_Cart_item+
		` WHERE `+
		user_id+
		` = ? `, userId).Iter()
	for iter.Scan(&cartItem.CartItemId, &cartItem.UserId,
		&cartItem.ProductId, &cartItem.Quantity,
		&cartItem.CreatedAt,
		&cartItem.ModifiedAt) {
		productIds = append(productIds, cartItem.ProductId)
		cartItemsArray = append(cartItemsArray, cartItem)
	}
	if err := iter.Close(); err != nil {
		log.Fatal("Address Error :: ", err)
	}

	// Reading products from Product Table
	for i, v := range productIds {
		var product = productTable.Read(session, v)
		cartItemsArray[i].Product = product
	}
	return cartItemsArray
}

func Delete(session *gocql.Session, userId string, cartItemId string) (bool, string) {
	err := session.Query(`Delete from`+
		` FROM `+table.TABLE_Cart_item+
		` WHERE `+
		user_id+
		`, `+
		cart_item_id+
		` = ? `, userId, cartItemId).Exec()

	if err != nil {
		fmt.Println("Error! cart_item table product delete :: ", err)
		return false, cartItemId
	}

	return true, cartItemId
}
