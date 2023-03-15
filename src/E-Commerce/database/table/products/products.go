package products

import (
	"E-Commerce/database/table"
	"fmt"
	"github.com/gocql/gocql"
	"time"
)

func CreateProduct(session *gocql.Session, sellerId string, sellerName string, catId string, subCatId string) {
	Insert(session,
		catId,
		"",
		"",
		"Methi Leaves",
		"https://www.bigbasket.com/media/uploads/p/s/20000974_12-fresho-methiventhaya-keerai-cleaned-without-roots.jpg",
		7.0,
		"Methi leaves are a very popular ingredient in Indian dishes that has a bitter taste and a strong aroma.",
		sellerId,
		sellerName,
		"",
		subCatId,
		0.100,
		1,
		true)
}

const (
	product_id      = "product_id"
	category_id     = "category_id"
	dimension       = "dimension"
	manufacturer_id = "manufacturer_id"
	name            = "name"
	picture         = "picture"
	price           = "price"
	description     = "description"
	seller_id       = "seller_id"
	seller_name     = "seller_name"
	sku             = "sku"
	sub_category_id = "sub_category_id"
	weight          = "weight"
	units           = "units"
	last_updated    = "last_updated"
	subscribable    = "subscribable"
)

type Product struct {
	ProductId      gocql.UUID `json:"product_id"`      //1
	CategoryId     string     `json:"category_id"`     //2
	Dimension      string     `json:"dimension"`       //3
	ManufacturerId string     `json:"manufacturer_id"` //4
	Name           string     `json:"name"`            //5
	Picture        string     `json:"picture"`         //6
	Price          float32    `json:"price"`           //7
	Description    string     `json:"description"`     //8
	SellerId       string     `json:"seller_id"`       //9
	SellerName     string     `json:"seller_name"`     //10
	Sku            string     `json:"sku"`             //11
	SubCategoryId  string     `json:"sub_category_id"` //12
	Weight         float32    `json:"weight"`          //13
	Units          int        `json:"units"`           //14
	LastUpdated    time.Time  `json:"last_Updated"`    //15
	Subscribable   bool       `json:"subscribable"`    //16
}

func Insert(session *gocql.Session, pCategoryId string, pDimension string, pManufacturerId string,
	pName string, pPicture string, pPrice float32, pDescription string, pSellerId string, pSellerName string, pSku string,
	pSubCategoryId string, pWeight float32, pUnits int, pSubscribable bool) bool {

	println("category_id", pCategoryId)
	println("dimension", pDimension)
	println("manufacturer_id", pManufacturerId)
	println("name", pName)
	println("picture", pPicture)
	println("price", pPrice)
	println("description", pDescription)
	println("seller_id", pSellerId)
	println("seller_name", pSellerName)
	println("sku", pSku)
	println("sub_category_id", pSubCategoryId)
	println("weight", pWeight)
	println("units", pUnits)
	println("subscribable", pSubscribable)

	if err := session.Query(`INSERT INTO `+table.TABLE_Products+`(
		`+product_id+`,		//1
		`+category_id+`,	//2
		`+dimension+`,		//3
		`+manufacturer_id+`,	//4
		`+name+`,			//5
		`+picture+`,		//6
		`+price+`,			//7
		`+description+`,	//8
		`+seller_id+`,		//9
		`+seller_name+`,	//10
		`+sku+`,			//11
		`+sub_category_id+`,//12
		`+weight+`,			//13
		`+units+`,			//14
		`+last_updated+`,	//15
		`+subscribable+`	//16
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?)`,
		gocql.TimeUUID(),
		pCategoryId,
		pDimension,
		pManufacturerId,
		pName,
		pPicture,
		pPrice,
		pDescription,
		pSellerId,
		pSellerName,
		pSku,
		pSubCategoryId,
		pWeight,
		pUnits,
		time.Now(),
		pSubscribable,
	).Exec(); err != nil {
		fmt.Println("Error! insert into Product ::::    ", err)
		return false
	}
	return true
}

func Read(session *gocql.Session, productId string) Product {
	var product = Product{}

	iter := session.Query(`SELECT `+
		product_id+
		`,`+category_id+
		`,`+dimension+
		`,`+manufacturer_id+
		`,`+name+
		`,`+picture+
		`,`+price+
		`,`+description+
		`,`+seller_id+
		`,`+seller_name+
		`,`+sku+
		`,`+sub_category_id+
		`,`+weight+
		`,`+units+
		`,`+last_updated+
		`,`+subscribable+
		` FROM `+table.TABLE_Products+
		` WHERE `+
		product_id+
		` = ? `, productId).Iter()
	for iter.Scan(&product.ProductId, &product.CategoryId,
		&product.Dimension, &product.ManufacturerId, &product.Name,
		&product.Picture, &product.Price, &product.Description,
		&product.SellerId, &product.SellerName, &product.Sku,
		&product.SubCategoryId, &product.Weight, &product.Units,
		&product.LastUpdated, &product.Subscribable) {
	}
	if err := iter.Close(); err != nil {
		fmt.Println("Error! Reading Product :: ", err)
		return Product{}
	}
	return product
}
