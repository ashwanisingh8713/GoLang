package category

import (
	"ServerDrivenUI/src/cassandra/db/table"
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"time"
)

const (
	category_id = "category_id"
	name        = "name"
	description = "description"
	created_at  = "created_at"
	modified_at = "modified_at"
	deleted_at  = "deleted_at"
	is_deleted  = "is_deleted"
)

type Category struct {
	CategoryId  string
	Name        string
	Description string
	CreatedAt   time.Time
	ModifiedAt  time.Time
	DeletedAt   time.Time
	IsDeleted   bool
}

func Insert(session *gocql.Session, catName string, catDescription string) {
	var category = Category{}
	category.Name = catName
	category.Description = catDescription
	category.CreatedAt = time.Now()
	category.ModifiedAt = time.Now()
	category.IsDeleted = false

	if err := session.Query(`INSERT INTO `+table.TABLE_Category+`(
		`+category_id+`,
		`+name+`,
		`+description+`,
		`+created_at+`,
		`+modified_at+`,
		`+is_deleted+`
		) VALUES (?, ?, ?, ?, ?, ?)`,
		gocql.TimeUUID(),
		category.Name,
		category.Description,
		category.CreatedAt,
		category.ModifiedAt,
		category.IsDeleted,
	).Exec(); err != nil {
		log.Fatal("Error! insert into Category ::::    ", err)
	}
}

func Read(session gocql.Session) {
	var category = Category{}
	iter := session.Query(`SELECT 
    			* FROM ` + table.TABLE_Category).Iter()

	for iter.Scan(&category.CategoryId, &category.IsDeleted, &category.CreatedAt, &category.DeletedAt,
		&category.Description,
		&category.ModifiedAt, &category.Name) {
		fmt.Println("Address : ", category.CategoryId, category.Name)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}
