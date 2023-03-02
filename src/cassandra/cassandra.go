package main

import (
	"ServerDrivenUI/src/cassandra/db/table/category"
	"ServerDrivenUI/src/cassandra/db/table/products"
	"ServerDrivenUI/src/cassandra/db/table/seller"
	"ServerDrivenUI/src/cassandra/db/table/sub_category"
)

var sellerIdFresho = "68c610a4-b8d8-11ed-88b6-16ab030fa2fa"
var sellerNameFresho = "Fresho"
var catIdVeggies = "cc9df97a-b8d3-11ed-9d8f-16ab030fa2fa"
var subCatIdLeafyVegetables = "68fe3078-b8d4-11ed-8c00-16ab030fa2fa"

func TEST() {

}

func main() {
	createDBSession()
	sellerInsertQuery()
	//subCategoryInsertQuery()
	//subCategoryInsertQuery()
	products.ProductInsertQuery(dbInstance, sellerIdFresho, sellerNameFresho, catIdVeggies, subCatIdLeafyVegetables)
	//categoryInsertQuery()
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
