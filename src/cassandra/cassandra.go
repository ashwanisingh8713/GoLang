package main

import (
	"ServerDrivenUI/src/cassandra/db/table/address"
	"ServerDrivenUI/src/cassandra/db/table/cart_item"
	"ServerDrivenUI/src/cassandra/db/table/category"
	"ServerDrivenUI/src/cassandra/db/table/seller"
	"ServerDrivenUI/src/cassandra/db/table/sub_category"
	"ServerDrivenUI/src/cassandra/db/table/user"
)

var sellerIdFresho = "61a58bb6-b8eb-11ed-9dee-16ab030fa2fa"
var sellerNameFresho = "Fresho"
var catIdVeggies = "580d281a-b8ec-11ed-bdf4-16ab030fa2fa"
var subCatIdLeafyVegetables = "73cb51bc-b8ec-11ed-a60c-16ab030fa2fa"
var userId = "97a6c53e-b9a6-11ed-bb41-16ab030fa2fa"

func TEST() {

}

func main() {
	createDBSession()
	//sellerInsertQuery()
	//categoryInsertQuery()
	//subCategoryInsertQuery()
	//products.ProductInsertQuery(dbInstance, sellerIdFresho, sellerNameFresho, catIdVeggies, subCatIdLeafyVegetables)
	//userInsertQuery()
	//cartItemInsertQuery()
	//addressInsertQuery()
	user.Read(dbInstance)
}

func addressInsertQuery() {
	var addressStruct = address.Address{}
	addressStruct.UserId = userId
	addressStruct.AddressType = "Office"
	addressStruct.AddressLine1 = "Ninestars Info"
	addressStruct.AddressLine2 = "JP Nagar 3rd Phase"
	addressStruct.Mobile1 = "8390098899"
	addressStruct.Mobile2 = "8390098890"
	addressStruct.Country = "India"
	addressStruct.City = "Bangalore"
	addressStruct.Zip = "560078"
	addressStruct.State = "Karnataka"
	addressStruct.IsPreferred = false

	address.Insert(dbInstance, addressStruct)
}

func cartItemInsertQuery() {
	cart_item.Insert(dbInstance, userId,
		"5fd267a0-b8f0-11ed-94a8-16ab030fa2fa",
		2)
}

func userInsertQuery() {
	user.Insert(dbInstance)
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
