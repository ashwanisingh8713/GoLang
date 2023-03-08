package cart_item

import (
	"E-Commerce/database/table"
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
	cartItemId gocql.UUID
	userId     string
	productId  string
	quantity   int
	createdAt  time.Time
	modifiedAt time.Time
}

func Insert(session *gocql.Session, userId string, productId string, productQuantity int) {
	var user = CartItem{}
	user.userId = userId
	user.productId = productId
	user.quantity = productQuantity
	user.createdAt = time.Now()
	user.modifiedAt = time.Now()

	if err := session.Query(`INSERT INTO `+table.TABLE_Cart_item+`(
		`+cart_item_id+`,
		`+user_id+`,
		`+product_id+`,
		`+quantity+`,
		`+created_at+`,
		`+modified_at+`
		) VALUES (?, ?, ?, ?, ?, ?)`,
		gocql.TimeUUID(),
		user.userId,
		user.productId,
		user.quantity,
		user.createdAt,
		user.modifiedAt,
	).Exec(); err != nil {
		log.Fatal("Error! insert into CartItem ::::    ", err)
	}
}
