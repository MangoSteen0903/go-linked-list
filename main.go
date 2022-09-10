package main

import (
	"github.com/MangoSteen0903/go-linkedlist/nodelist"
)

func main() {
	newList := nodelist.LinkedList()
	newList.Push("First")
	newList.Push("Second")
	newList.Push("Third")
	newList.Push("Fourth")

	newList.ReverseList()
	newList.ViewList()
}
