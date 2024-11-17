package main

import (
	"container/list"
	"fmt"
)

func main() {
	// var data list.List = list.List{}
	var data *list.List = list.New()

	data.PushBack("Geralt")
	data.PushBack("Yennefer")
	data.PushBack("Triss")

	var head *list.Element = data.Front()
	fmt.Println("Head: ", head.Value)

	next := head.Next()
	fmt.Println("Next: ", next.Value)

	next = next.Next()
	fmt.Println("Next: ", next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
