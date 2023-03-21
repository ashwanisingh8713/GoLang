package sub_category

import (
	"E-Commerce/database/table"
	"fmt"
	"github.com/gocql/gocql"
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

func CreateSubCategory(session *gocql.Session, categoryId string, subCatName string, subCatDescription string) (bool, string) {
	var subCategory = SubCategory{}
	subCategory.Name = subCatName
	subCategory.Description = subCatDescription
	subCategory.CreatedAt = time.Now()
	subCategory.ModifiedAt = time.Now()
	uniqueId := gocql.TimeUUID()

	if err := session.Query(` INSERT INTO `+table.TABLE_Sub_category+`(
		`+sub_category_id+`,
		`+category_id+`,
		`+name+`,
		`+description+`,
		`+created_at+`,
		`+modified_at+`,
		`+picture+`
		) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		uniqueId,
		categoryId,
		subCategory.Name,
		subCategory.Description,
		subCategory.CreatedAt,
		subCategory.ModifiedAt,
		subCategory.Picture,
	).Exec(); err != nil {
		fmt.Println("Error! insert into Sub-Category ::::    ", err)
		return false, "Failed to create Sub-Category"
	}
	return true, uniqueId.String()
}

func GetAllSubCategory(session *gocql.Session) (bool, []SubCategory) {
	var subCategories []SubCategory
	var subCategory = SubCategory{}
	iter := session.Query(`SELECT 
    			` + sub_category_id + `,
		` + category_id + `,
		` + name + `,
		` + description + `,
		` + created_at + `,
		` + modified_at + `,
		` + picture + `
				FROM ` + table.TABLE_Sub_category).Iter()

	for iter.Scan(&subCategory.SubCategoryId, &subCategory.CategoryId, &subCategory.Name, &subCategory.Description,
		&subCategory.CreatedAt, &subCategory.ModifiedAt, &subCategory.Picture) {
		subCategories = append(subCategories, subCategory)
	}
	if err := iter.Close(); err != nil {
		return false, []SubCategory{}
	}
	return true, subCategories
}

func GetSubCategoryById(session *gocql.Session, subCategoryId string) (bool, []SubCategory) {
	subCategories := []SubCategory{}
	var subCategory = SubCategory{}
	iter := session.Query(`SELECT 
		`+sub_category_id+
		`,`+category_id+
		`,`+name+
		`,`+description+
		`,`+created_at+
		`,`+modified_at+
		`,`+picture+
		` FROM `+table.TABLE_Sub_category+
		` WHERE `+
		sub_category_id+
		` = ? `, subCategoryId).Iter()

	for iter.Scan(&subCategory.SubCategoryId, &subCategory.CategoryId,
		&subCategory.Name, &subCategory.Description, &subCategory.CreatedAt, &subCategory.ModifiedAt, &subCategory.Picture) {
		subCategories = append(subCategories, subCategory)
	}
	if err := iter.Close(); err != nil {
		return false, []SubCategory{}
	}
	return true, subCategories
}
