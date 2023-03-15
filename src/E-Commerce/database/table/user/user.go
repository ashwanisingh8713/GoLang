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

func CreateUser(session *gocql.Session, u_firstName, u_lastName, u_mobile, u_email, u_password, u_loginMode string) bool {
	var status = true
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
		u_firstName,
		u_lastName,
		u_mobile,
		u_email,
		u_password,
		u_loginMode,
		time.Now(),
	).Exec(); err != nil {
		fmt.Println("Error! insert into User ::::    ", err)
		status = false
	}
	return status
}

func GetUserInfo(session *gocql.Session, userId string) User {
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
		fmt.Println("User Table, Get User Info failed error is ", err)
	}
	return user
}

func GetAllUser(session *gocql.Session) ([]User, bool) {
	var status = true
	var users []User
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
		users = append(users, user)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
		status = false
	}
	return users, status
}
