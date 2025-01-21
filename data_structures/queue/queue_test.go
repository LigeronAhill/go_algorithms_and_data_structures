package queue

import "testing"

func TestQueue(t *testing.T) {
	queue := NewQueue()
	order1 := NewOrder("Greg White")
	order1.Priority(2)
	order1.Quantity(20)
	order1.Product("Computer")
	order2 := NewOrder("John Smith")
	order2.Priority(1)
	order2.Quantity(10)
	order2.Product("Monitor")
	queue.Add(order1)
	queue.Add(order2)
	for i := range *queue {
		t.Logf("%+v\n", (*queue)[i])
	}
}
