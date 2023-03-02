package main

import (
	"ServerDrivenUI/src/cassandra/db/table/category"
	"ServerDrivenUI/src/cassandra/db/table/products"
	"ServerDrivenUI/src/cassandra/db/table/seller"
	"ServerDrivenUI/src/cassandra/db/table/sub_category"
)

var sellerIdFresho = "61a58bb6-b8eb-11ed-9dee-16ab030fa2fa"
var sellerNameFresho = "Fresho"
var catIdVeggies = "580d281a-b8ec-11ed-bdf4-16ab030fa2fa"
var subCatIdLeafyVegetables = "73cb51bc-b8ec-11ed-a60c-16ab030fa2fa"

func TEST() {

}

func main() {
	createDBSession()
	//sellerInsertQuery()
	//categoryInsertQuery()
	//subCategoryInsertQuery()
	products.ProductInsertQuery(dbInstance, sellerIdFresho, sellerNameFresho, catIdVeggies, subCatIdLeafyVegetables)
}

func categoryInsertQuery() {
	category.Insert(dbInstance,
		"Veggies",
		"It is Vegetable category",
		"",
	)
}

func subCategoryInsertQuery() {
	sub_category.Insert(dbInstance,
		catIdVeggies,
		"Leafy Vegetables",
		"Leaf vegetables, also called leafy greens, pot herbs, vegetable greens, or simply greens, are plant leaves eaten as a vegetable, sometimes accompanied by tender petioles and shoots. ")
}

func sellerInsertQuery() {
	seller.Insert(dbInstance, "Fresho")
}
