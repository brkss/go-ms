package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

func main() {

	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection Error : ", err)
	}

	a := Item{"client item 1", "client item 1 body"}
	//b := Item {"client item 2", "client item 2 body"}
	//c := Item {"client item 3", "client item 3 body"}

	client.Call("API.AddItem", a, &reply)
	client.Call("API.GetItems", "", &db)

	fmt.Println("database from client : ", db)

}
