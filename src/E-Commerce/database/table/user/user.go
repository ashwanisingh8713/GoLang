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

func Insert(session *gocql.Session) {
	var user = User{}
	user.FirstName = "User - 1 "
	user.LastName = "Dev Team"
	user.Mobile = "8390089995"
	user.Email = "ashwani@gmail.com"
	user.Password = "Password"
	user.LoginMode = "Google"
	user.CreatedAt = time.Now()

	if err := session.Query(`INSERT INTO `+table.TABLE_User+`(
		`+user_id+`,
		`+first_name+`,
		`+last_name+`,
		`+mobile+`,
		`+email+`,
		`+password+`,
		`+login_mode+`,
		`+created_at+`
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		gocql.TimeUUID(),
		user.FirstName,
		user.LastName,
		user.Mobile,
		user.Email,
		user.Password,
		user.LoginMode,
		user.CreatedAt,
	).Exec(); err != nil {
		log.Fatal("Error! insert into User ::::    ", err)
	}
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
