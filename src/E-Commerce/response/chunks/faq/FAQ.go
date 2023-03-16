package faq

import (
	"E-Commerce/data"
)

func Faq() interface{} {
	//var viewGrpChildren []data.FAQ

	var queryOrderList []data.Query
	var queryPurchaseList []data.Query
	var querySellList []data.Query
	var queryList []data.Faqs

	var queryOrder1 = data.Query{
		Question: string("Where do you Deliver?"),
		Answer:   string("Placing an order is very simple.  Just register on the SG Grocery website/mobile application, pick your choice of products with a wide range of product selection in the online store and proceed to checkout or just call customer care and place an order. i.e. 1800-123-1947 "),
	}
	var queryOrder2 = data.Query{
		Question: string("How can I order at SG Grocery?"),
		Answer:   string("Placing an order is very simple.  Just register on the SG Grocery website/mobile application, pick your choice of products with a wide range of product selection in the online store and proceed to checkout or just call customer care and place an order. i.e. 1800-123-1947 "),
	}

	var queryOrder3 = data.Query{
		Question: string("How do I know at what time I will receive my Delivery?"),
		Answer:   string("Placing an order is very simple.  Just register on the SG Grocery website/mobile application, pick your choice of products with a wide range of product selection in the online store and proceed to checkout or just call customer care and place an order. i.e. 1800-123-1947 "),
	}

	var queryOrder4 = data.Query{
		Question: string("What is minimum order value?"),
		Answer:   string("Placing an order is very simple.  Just register on the SG Grocery website/mobile application, pick your choice of products with a wide range of product selection in the online store and proceed to checkout or just call customer care and place an order. i.e. 1800-123-1947 "),
	}

	var queryOrder5 = data.Query{
		Question: string("What is minimum order value?"),
		Answer:   string("Placing an order is very simple.  Just register on the SG Grocery website/mobile application, pick your choice of products with a wide range of product selection in the online store and proceed to checkout or just call customer care and place an order. i.e. 1800-123-1947 "),
	}

	var queryOrder6 = data.Query{
		Question: string("What if I want to return something?"),
		Answer:   string("Placing an order is very simple.  Just register on the SG Grocery website/mobile application, pick your choice of products with a wide range of product selection in the online store and proceed to checkout or just call customer care and place an order. i.e. 1800-123-1947 "),
	}
	queryOrderList = append(queryOrderList, queryOrder1)
	queryOrderList = append(queryOrderList, queryOrder2)
	queryOrderList = append(queryOrderList, queryOrder3)
	queryOrderList = append(queryOrderList, queryOrder4)
	queryOrderList = append(queryOrderList, queryOrder5)
	queryOrderList = append(queryOrderList, queryOrder6)

	var queryPurchase1 = data.Query{
		Question: string("Where do you purchase Deliver?"),
		Answer:   string("Placing an order is very simple.  Just register on the SG Grocery website/mobile application, pick your choice of products with a wide range of product selection in the online store and proceed to checkout or just call customer care and place an order. i.e. 1800-123-1947 "),
	}
	var queryPurchase2 = data.Query{
		Question: string("How can I purchase at SG Grocery?"),
		Answer:   string("Placing an order is very simple.  Just register on the SG Grocery website/mobile application, pick your choice of products with a wide range of product selection in the online store and proceed to checkout or just call customer care and place an order. i.e. 1800-123-1947 "),
	}

	var queryPurchase3 = data.Query{
		Question: string("I will receive my purchased  Delivery?"),
		Answer:   string("Placing an order is very simple.  Just register on the SG Grocery website/mobile application, pick your choice of products with a wide range of product selection in the online store and proceed to checkout or just call customer care and place an order. i.e. 1800-123-1947 "),
	}

	var queryPurchase4 = data.Query{
		Question: string("What is minimum purchase order value?"),
		Answer:   string("Placing an order is very simple.  Just register on the SG Grocery website/mobile application, pick your choice of products with a wide range of product selection in the online store and proceed to checkout or just call customer care and place an order. i.e. 1800-123-1947 "),
	}

	var queryPurchase5 = data.Query{
		Question: string("What is minimum purchase value?"),
		Answer:   string("Placing an order is very simple.  Just register on the SG Grocery website/mobile application, pick your choice of products with a wide range of product selection in the online store and proceed to checkout or just call customer care and place an order. i.e. 1800-123-1947 "),
	}

	var queryPurchase6 = data.Query{
		Question: string("What if I want to return purchased?"),
		Answer:   string("Placing an order is very simple.  Just register on the SG Grocery website/mobile application, pick your choice of products with a wide range of product selection in the online store and proceed to checkout or just call customer care and place an order. i.e. 1800-123-1947 "),
	}

	queryPurchaseList = append(queryPurchaseList, queryPurchase1)
	queryPurchaseList = append(queryPurchaseList, queryPurchase2)
	queryPurchaseList = append(queryPurchaseList, queryPurchase3)
	queryPurchaseList = append(queryPurchaseList, queryPurchase4)
	queryPurchaseList = append(queryPurchaseList, queryPurchase5)
	queryPurchaseList = append(queryPurchaseList, queryPurchase6)

	var querySell1 = data.Query{
		Question: string("Where do you sell Deliver?"),
		Answer:   string("Placing an order is very simple.  Just register on the SG Grocery website/mobile application, pick your choice of products with a wide range of product selection in the online store and proceed to checkout or just call customer care and place an order. i.e. 1800-123-1947 "),
	}
	var querySell2 = data.Query{
		Question: string("How can I sell at SG Grocery?"),
		Answer:   string("Placing an order is very simple.  Just register on the SG Grocery website/mobile application, pick your choice of products with a wide range of product selection in the online store and proceed to checkout or just call customer care and place an order. i.e. 1800-123-1947 "),
	}

	var querySell3 = data.Query{
		Question: string("I will receive my sell  Delivery?"),
		Answer:   string("Placing an order is very simple.  Just register on the SG Grocery website/mobile application, pick your choice of products with a wide range of product selection in the online store and proceed to checkout or just call customer care and place an order. i.e. 1800-123-1947 "),
	}

	var querySell4 = data.Query{
		Question: string("What is minimum sell order value?"),
		Answer:   string("Placing an order is very simple.  Just register on the SG Grocery website/mobile application, pick your choice of products with a wide range of product selection in the online store and proceed to checkout or just call customer care and place an order. i.e. 1800-123-1947 "),
	}

	var querySell5 = data.Query{
		Question: string("What is minimum sell value?"),
		Answer:   string("Placing an order is very simple.  Just register on the SG Grocery website/mobile application, pick your choice of products with a wide range of product selection in the online store and proceed to checkout or just call customer care and place an order. i.e. 1800-123-1947 "),
	}

	var querySell6 = data.Query{
		Question: string("What if I want to return sell?"),
		Answer:   string("Placing an order is very simple.  Just register on the SG Grocery website/mobile application, pick your choice of products with a wide range of product selection in the online store and proceed to checkout or just call customer care and place an order. i.e. 1800-123-1947 "),
	}

	querySellList = append(querySellList, querySell1)
	querySellList = append(querySellList, querySell2)
	querySellList = append(querySellList, querySell3)
	querySellList = append(querySellList, querySell4)
	querySellList = append(querySellList, querySell5)
	querySellList = append(querySellList, querySell6)

	var query1 = data.Faqs{
		Title:   string("Order"),
		Queries: queryOrderList,
	}
	var query2 = data.Faqs{
		Title:   string("Purchase"),
		Queries: queryPurchaseList,
	}
	var query3 = data.Faqs{
		Title:   string("Sell"),
		Queries: queryPurchaseList,
	}

	queryList = append(queryList, query1)
	queryList = append(queryList, query2)
	queryList = append(queryList, query3)

	return queryList
}
