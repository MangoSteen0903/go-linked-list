package nodelist

type Node struct {
	Data string
	Next *Node
}

func createNode(data string) *Node {
	newNode := &Node{
		Data: data,
		Next: &Node{Data: ""},
	}
	return newNode
}
