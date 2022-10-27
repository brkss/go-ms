package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type API int

type Item struct {
	Title string
	Body  string
}

var database []Item

func (a *API) GetItems(title string, reply *[]Item) error {
	*reply = database
	return nil
}

func (a *API) GetByName(title string, reply *Item) error {
	var getItem Item

	for _, val := range database {
		if val.Title == title {
			getItem = val
			break
		}
	}
	*reply = getItem
	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

func (a *API) EditItem(title string, edit Item, reply *Item) error {
	var changed Item

	for i, val := range database {
		if val.Title == edit.Title {
			database[i] = edit
			changed = edit
			break
		}
	}
	*reply = changed
	return nil
}

func (a *API) DeleteItem(del Item, reply *Item) error {
	var deleted Item

	for index, val := range database {
		if val.Title == del.Title && val.Body == del.Body {
			database = append(database[:index], database[index+1:]...)
			break
		}
	}
	*reply = deleted
	return nil
}

func main() {

	var api = new(API)
	err := rpc.Register(api)

	if err != nil {
		log.Fatal("Error registering API : ", err)
	}

	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Listener error : ", err)
	}

	log.Println("🚀 server running at http://localhost:40404")
	err = http.Serve(listener, nil)

	if err != nil {
		log.Fatal("Error serving : ", err)
	}

	/*
		fmt.Println("database : ", database)
		a := Item{"item 1", "this item 1 body"}
		b := Item{"item 2", "this item 2 body"}
		c := Item{"item 3", "this item 3 body"}

		AddItem(a)
		AddItem(b)
		AddItem(c)

		fmt.Println("database : ", database)
		DeleteItem(b)
		fmt.Println("database (delete item 2) : ", database)

		EditItem("item 1", Item{"item 1 edited", "this item 1 body (edited)"})
		fmt.Println("database (item 1 edited) : ", database)
	*/
}
