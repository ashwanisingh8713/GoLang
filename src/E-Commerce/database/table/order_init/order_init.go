package order_init

import (
	"github.com/gocql/gocql"
	"time"
)

const (
	order_id = "order_id"
	user_id  = "user_id"
	mobile   = "mobile"
	email_id = "email_id"
	order_at = "order_at"
)

type OrderInit struct {
	order_id gocql.UUID
	user_id  string
	mobile   string
	email_id string
	order_at time.Time
}

func createOrderInit(session *gocql.Session, userId string, mobile string, emailId string) {

}
