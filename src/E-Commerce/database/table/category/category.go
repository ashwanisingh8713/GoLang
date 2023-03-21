package category

import (
	"E-Commerce/database/table"
	"fmt"
	"github.com/gocql/gocql"
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
	picture     = "picture"
)

type Category struct {
	CategoryId  gocql.UUID
	Name        string
	Description string
	CreatedAt   time.Time
	ModifiedAt  time.Time
	DeletedAt   time.Time
	IsDeleted   bool
	Picture     string
}

func CreateCategory(session *gocql.Session, catName string, catDescription string, catPicture string) {
	/*var category = Category{}
	category.Name = catName
	category.Description = catDescription
	category.CreatedAt = time.Now()
	category.ModifiedAt = time.Now()
	category.IsDeleted = false*/

	if err := session.Query(`INSERT INTO `+table.TABLE_Category+`(
		`+category_id+`,
		`+name+`,
		`+description+`,
		`+created_at+`,
		`+modified_at+`,
		`+is_deleted+`,
		`+picture+`
		) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		gocql.TimeUUID(),
		catName,
		catDescription,
		time.Now(),
		time.Now(),
		true,
		catPicture,
	).Exec(); err != nil {
		fmt.Println("Error! insert into Category ::::    ", err)
	}
}

func GetAllCategory(session *gocql.Session) (bool, []Category) {
	var categories []Category
	var category = Category{}
	iter := session.Query(`SELECT 
    			` + category_id + `,
		` + name + `,
		` + description + `,
		` + created_at + `,
		` + modified_at + `,
		` + is_deleted + `,
		` + picture + `
				FROM ` + table.TABLE_Category).Iter()

	for iter.Scan(&category.CategoryId, &category.Name, &category.Description, &category.CreatedAt,
		&category.ModifiedAt,
		&category.IsDeleted, &category.Picture) {
		categories = append(categories, category)
	}
	if err := iter.Close(); err != nil {
		return false, []Category{}
	}
	return true, categories
}

func GetCategoryById(session *gocql.Session, categoryId string) (bool, Category) {
	var category = Category{}
	err := session.Query(`SELECT `+
		category_id+
		`,`+name+
		`,`+description+
		`,`+created_at+
		`,`+modified_at+
		`,`+deleted_at+
		`,`+is_deleted+
		`,`+picture+
		` FROM `+table.TABLE_Category+
		` WHERE `+
		category_id+
		` = ? `, categoryId).Consistency(gocql.One).Scan(&category.CategoryId, &category.Name, &category.Description, &category.CreatedAt,
		&category.ModifiedAt,
		&category.DeletedAt, &category.IsDeleted, &category.Picture)
	if err != nil {
		fmt.Println("Category Table, Get Category   Info failed error is ", err)
		return false, category
	}
	return true, category
}
