package nodelist

import (
	"fmt"
	"sync"
)

type linkedlist struct {
	Head   *Node
	Tail   *Node
	Height int
}

var l *linkedlist
var once sync.Once

func (l *linkedlist) Push(data string) *Node {
	newNode := createNode(data)
	if l.Tail != nil {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
	l.Height++
	return newNode
}

func (l *linkedlist) Pop() *Node {
	tailNode := l.Tail
	prevNode := l.FindNode(l.Height - 1)
	prevNode.Next = &Node{Data: ""}
	l.Tail = prevNode
	l.Height--
	return tailNode
}

func (l *linkedlist) Insert(data string, height int) {
	if height == l.Height {
		l.Push(data)
	} else {
		newNode := createNode(data)
		nextNode := l.FindNode(height)
		prevNode := l.FindNode(height - 1)

		prevNode.Next = newNode
		newNode.Next = nextNode

		l.Height++
	}
}

func (l *linkedlist) Delete(height int) {
	if height == l.Height {
		l.Pop()
	} else {
		prevNode := l.FindNode(height - 1)
		nextNode := l.FindNode(height + 1)
		prevNode.Next = nextNode
		l.Height--
	}
}

func (l *linkedlist) FindNode(height int) *Node {
	currentNode := l.Head
	for i := 0; i < height-1; i++ {
		currentNode = currentNode.Next
	}
	return currentNode
}

func (l *linkedlist) ViewHeight() {
	fmt.Println(l.Height)
}

func (l *linkedlist) ViewList() {
	currentNode := l.Head
	for {
		if currentNode.Next.Data == "" {
			fmt.Printf("Data: %s\n", currentNode.Data)
			fmt.Printf("%s is a Tail Node.\n\n", currentNode.Data)
			break
		} else {
			fmt.Printf("Data: %s\n", currentNode.Data)
			fmt.Printf("Next Data: %s\n\n", currentNode.Next.Data)
			currentNode = currentNode.Next
		}
	}
}

func (l *linkedlist) ReverseList() {
	currentNode := l.Tail
	currentHeight := l.Height
	for {
		if l.Head.Data == currentNode.Data {
			l.Head = l.Tail
			l.Tail = currentNode
			currentNode.Next = &Node{Data: ""}
			break
		}
		prevNode := l.FindNode(currentHeight - 1)
		currentNode.Next = prevNode
		currentNode = prevNode
		currentHeight--
	}
}

func LinkedList() *linkedlist {
	if l == nil {
		once.Do(func() {
			l = &linkedlist{}
			node := l.Push("Genesis")
			l.Head = node
			l.Tail = node
		})
	}
	return l
}
