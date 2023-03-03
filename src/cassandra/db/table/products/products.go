package products

import (
	"ServerDrivenUI/src/cassandra/db/table"
	"github.com/gocql/gocql"
	"log"
	"time"
)

func ProductInsertQuery(session *gocql.Session, sellerId string, sellerName string, catId string, subCatId string) {
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
	ProductId      gocql.UUID //1
	CategoryId     string     //2
	Dimension      string     //3
	ManufacturerId string     //4
	Name           string     //5
	Picture        string     //6
	Price          float32    //7
	Description    string     //8
	SellerId       string     //9
	SellerName     string     //10
	Sku            string     //11
	SubCategoryId  string     //12
	Weight         float32    //13
	Units          int        //14
	LastUpdated    time.Time  //15
	Subscribable   bool       //16
}

func Insert(session *gocql.Session, pCategoryId string, pDimension string, pManufacturerId string,
	pName string, pPicture string, pPrice float32, pDescription string, pSellerId string, pSellerName string, pSku string,
	pSubCategoryId string, pWeight float32, pUnits int, pSubscribable bool) {

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
		log.Fatal("Error! insert into Product ::::    ", err)
	}
}
