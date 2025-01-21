package doublelinkedlist

import (
	"fmt"
	"testing"
)

func TestDLL(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 7)
	fmt.Println(linkedList.headNode.property)
	node := linkedList.NodeBetweenValues(1, 5)
	fmt.Println(node.property)
}
