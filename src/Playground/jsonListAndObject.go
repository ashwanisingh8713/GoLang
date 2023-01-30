package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var clients []Client
	clients = append(clients, Client{
		Hostname: "localhost",
		IP:       "127.0.0.1",
		MacAddr:  "1123:22512:25632",
	}, Client{
		Hostname: "localhost222",
		IP:       "122227.0.0.1",
		MacAddr:  "00001123:22512:25632",
	})

	// add more if you want ...

	myJson, _ := json.Marshal(Connection{Clients: clients})
	fmt.Println(string(myJson))
}

type Client struct {
	Hostname string `json:"Hostname"`
	IP       string `json:"IP"`
	MacAddr  string `json:"MacAddr"`
}

type Connection struct {
	Clients []Client `json:"Clients"`
}
