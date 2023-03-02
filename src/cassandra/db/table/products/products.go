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
		"Coriander Leaves",
		"https://www.bigbasket.com/media/uploads/p/s/10000097_19-fresho-coriander-leaves.jpg",
		0.9000,
		"Coriander leaves are green, fragile with a decorative appearance. They contain minimal aroma and have a spicy sweet taste.",
		sellerId,
		sellerName,
		"",
		subCatId,
		0.2500,
		1,
		false)
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
	ProductId      gocql.UUID
	CategoryId     string
	Dimension      string
	ManufacturerId string
	Name           string
	Picture        string
	Price          float32
	Description    string
	SellerId       string
	SellerName     string
	Sku            string
	SubCategoryId  string
	Weight         float32
	Units          int
	LastUpdated    time.Time
	Subscribable   bool
}

func Insert(session *gocql.Session, pCategoryId string, pDimension string, pManufacturerId string,
	pName string, pPicture string, pPrice float32, pDescription string, pSellerId string, pSellerName string, pSku string,
	pSubCategoryId string, pWeight float32, pUnits int, pSubscribable bool) {

	if err := session.Query(`INSERT INTO `+table.TABLE_Products+`(
		`+product_id+`,
		`+category_id+`,
		`+dimension+`,
		`+manufacturer_id+`,
		`+name+`,
		`+picture+`,
		`+price+`,
		`+description+`,
		`+seller_id+`,
		`+seller_name+`,
		`+sku+`,
		`+sub_category_id+`,
		`+weight+`,
		`+units+`,
		`+last_updated+`,
		`+subscribable+`
		) VALUES (?, ?, ?, ?, ?, ?,?, ?, ?, ?, ?, ?,?,?,?,?)`,
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
