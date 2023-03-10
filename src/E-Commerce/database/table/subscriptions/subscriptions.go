package subscriptions

import (
	"E-Commerce/database/table"
	productTable "E-Commerce/database/table/products"
	"github.com/gocql/gocql"
	"log"
	"time"
)

const (
	subscription_id = "subscription_id"
	product_id      = "product_id"
	order_item_id   = "order_id"
	user_id         = "user_id"
	monday          = "monday"
	tuesday         = "tuesday"
	wednesday       = "wednesday"
	thursday        = "thursday"
	friday          = "friday"
	saturday        = "saturday"
	sunday          = "sunday"
	start_date      = "start_date"
	end_date        = "end_date"
	created_at      = "created_at"
	pause_dates     = "pause_dates"
)

type Subscription struct {
	SubscriptionId string               `json:"subscription_id"`
	ProductId      string               `json:"product_id"`
	OrderItemId    string               `json:"order_id"` // This is OrderItemId
	UserId         string               `json:"user_id"`
	Monday         bool                 `json:"monday"`
	Tuesday        bool                 `json:"tuesday"`
	Wednesday      bool                 `json:"wednesday"`
	Thursday       bool                 `json:"thursday"`
	Friday         bool                 `json:"friday"`
	Saturday       bool                 `json:"saturday"`
	Sunday         bool                 `json:"sunday"`
	StartDate      time.Time            `json:"startDate"`
	EndDate        time.Time            `json:"endDate"`
	CreatedAt      time.Time            `json:"createdAt"`
	PauseDates     []string             `json:"pauseDates"`
	Product        productTable.Product `json:"product,omitempty"`
}

func Insert(session *gocql.Session, userId string, productId string, orderItemId string,
	Monday bool, Tuesday bool, Wednesday bool, Thursday bool, Friday bool, Saturday bool, Sunday bool,
	StartDate time.Time, EndDate time.Time, PauseDates []string) {
	var subscription = Subscription{}
	subscription.ProductId = productId
	subscription.OrderItemId = orderItemId
	subscription.UserId = userId
	subscription.Monday = Monday
	subscription.Tuesday = Tuesday
	subscription.Wednesday = Wednesday
	subscription.Thursday = Thursday
	subscription.Friday = Friday
	subscription.Saturday = Saturday
	subscription.Sunday = Sunday
	subscription.StartDate = StartDate
	subscription.EndDate = EndDate
	subscription.CreatedAt = time.Now()
	subscription.PauseDates = PauseDates

	if err := session.Query(`INSERT INTO `+table.TABLE_Subscriptions+`(
		`+subscription_id+`,
		`+product_id+`,
		`+order_item_id+`,
		`+user_id+`,
		`+monday+`,
		`+tuesday+`,
		`+wednesday+`,
		`+thursday+`,
		`+friday+`,
		`+saturday+`,
		`+sunday+`,
		`+start_date+`,
		`+end_date+`,
		`+created_at+`,
		`+pause_dates+`
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		gocql.TimeUUID(),
		subscription.ProductId,
		subscription.OrderItemId,
		subscription.UserId,
		subscription.Monday,
		subscription.Tuesday,
		subscription.Wednesday,
		subscription.Thursday,
		subscription.Friday,
		subscription.Saturday,
		subscription.Sunday,
		subscription.StartDate,
		subscription.EndDate,
		subscription.CreatedAt,
		subscription.PauseDates,
	).Exec(); err != nil {
		log.Fatal("Error! insert into CartItem ::::    ", err)
	}
}

func ReadAllSubcriptions(session *gocql.Session, userId string) []Subscription {
	var subscriptionList []Subscription
	var productIds []string
	var subscription = Subscription{}

	iter := session.Query(`SELECT `+
		subscription_id+
		`,`+product_id+
		`,`+order_item_id+
		`,`+user_id+
		`,`+monday+
		`,`+tuesday+
		`,`+wednesday+
		`,`+thursday+
		`,`+friday+
		`,`+saturday+
		`,`+sunday+
		`,`+start_date+
		`,`+end_date+
		`,`+created_at+
		`,`+pause_dates+
		` FROM `+table.TABLE_Subscriptions+
		` WHERE `+
		user_id+
		` = ? `, userId).Iter()
	for iter.Scan(&subscription.SubscriptionId, &subscription.ProductId,
		&subscription.OrderItemId, &subscription.UserId,
		&subscription.Monday, &subscription.Tuesday, &subscription.Wednesday, &subscription.Thursday, &subscription.Friday, &subscription.Saturday, &subscription.Sunday,
		&subscription.StartDate, &subscription.EndDate, &subscription.CreatedAt, &subscription.PauseDates) {
		productIds = append(productIds, subscription.ProductId)
		subscriptionList = append(subscriptionList, subscription)
	}
	if err := iter.Close(); err != nil {
		log.Fatal("Address Error :: ", err)
	}

	// Reading products from Product Table
	for index, productId := range productIds {
		var product = productTable.Read(session, productId)
		subscriptionList[index].Product = product
	}
	return subscriptionList
}
