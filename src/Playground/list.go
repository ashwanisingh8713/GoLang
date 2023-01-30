package main

import "log"

type User struct {
	Name string
}

type Article struct {
	Title string
}

func main1() {
	users := []User{}
	articles := []Article{}

	user := User{Name: "Ashwani"}
	users = append(users, user)

	models := []interface{}{users, articles}
	for _, model := range models {
		log.Printf("%#v", model)
	}
}
