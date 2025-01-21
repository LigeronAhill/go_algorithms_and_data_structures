package queue

type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

func NewOrder(customerName string) *Order {
	return &Order{customerName: customerName}
}

func (o *Order) Priority(priority int) {
	o.priority = priority
}

func (o *Order) Quantity(quantity int) {
	o.quantity = quantity
}

func (o *Order) Product(product string) {
	o.product = product
}

type Queue []*Order

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Add(order *Order) {
	if len(*q) == 0 {
		*q = append(*q, order)
	} else {
		appended := false
		for i, addedOrder := range *q {
			if order.priority > addedOrder.priority {
				*q = append((*q)[:i], append(Queue{order}, (*q)[i:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*q = append(*q, order)
		}
	}
}
