package main

import (
	"ServerDrivenUI/src/cassandra/db/table/category"
)

func main() {
	getInstance()
	cassandraSetup()
}

func cassandraSetup() {
	category.Insert(dbInstance, "Groceries", "It is Groceries category.")
}
