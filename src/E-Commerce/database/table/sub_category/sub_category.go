package sub_category

import (
	"E-Commerce/database/table"
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"time"
)

const (
	sub_category_id = "sub_category_id"
	category_id     = "category_id"
	name            = "name"
	description     = "description"
	created_at      = "created_at"
	modified_at     = "modified_at"
	deleted_at      = "deleted_at"
	picture         = "picture"
)

type SubCategory struct {
	SubCategoryId gocql.UUID
	CategoryId    string
	Name          string
	Description   string
	CreatedAt     time.Time
	ModifiedAt    time.Time
	DeletedAt     time.Time
	Picture       string
}

func Insert(session *gocql.Session, categoryId string, subCatName string, subCatDescription string) {
	var subCategory = SubCategory{}
	subCategory.Name = subCatName
	subCategory.Description = subCatDescription
	subCategory.CreatedAt = time.Now()
	subCategory.ModifiedAt = time.Now()

	if err := session.Query(` INSERT INTO `+table.TABLE_Sub_category+`(
		`+sub_category_id+`,
		`+category_id+`,
		`+name+`,
		`+description+`,
		`+created_at+`,
		`+modified_at+`,
		`+picture+`
		) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		gocql.TimeUUID(),
		categoryId,
		subCategory.Name,
		subCategory.Description,
		subCategory.CreatedAt,
		subCategory.ModifiedAt,
		subCategory.Picture,
	).Exec(); err != nil {
		log.Fatal("Error! insert into Sub-Category ::::    ", err)
	}
}

func Read(session gocql.Session) {
	var subCategory = SubCategory{}
	iter := session.Query(`SELECT 
    			* FROM ` + table.TABLE_Sub_category).Iter()

	for iter.Scan(&subCategory.SubCategoryId, &subCategory.CategoryId, &subCategory.Name, &subCategory.CreatedAt,
		&subCategory.DeletedAt,
		&subCategory.Description,
		&subCategory.ModifiedAt) {
		fmt.Println("Address : ", subCategory.CategoryId, subCategory.Name)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}
