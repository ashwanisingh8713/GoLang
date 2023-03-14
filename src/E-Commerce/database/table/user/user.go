package user

import (
	"E-Commerce/database/table"
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"time"
)

type User struct {
	UserId    string    `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Mobile    string    `json:"mobile"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	LoginMode string    `json:"login_mode"`
	CreatedAt time.Time `json:"created_at"`
}

const (
	user_id    = "user_id"
	first_name = "first_name"
	last_name  = "last_name"
	mobile     = "mobile"
	email      = "email"
	password   = "password"
	login_mode = "login_mode"
	created_at = "created_at"
)

func Read(session *gocql.Session, userId string) User {
	var user = User{}
	iter := session.Query(`SELECT ` +
		user_id +
		`,` + first_name +
		`,` + last_name +
		`,` + mobile +
		`,` + email +
		`,` + password +
		`,` + login_mode +
		`,` + created_at +
		` FROM ` + table.TABLE_User).Iter()

	for iter.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.Mobile,
		&user.Email,
		&user.Password, &user.LoginMode, &user.CreatedAt) {
		fmt.Println("User : ", user.FirstName, user.UserId)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return user
}
