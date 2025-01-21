package heaps

import (
	"container/heap"
	"fmt"
)

type integerHeap []int

func (iheap integerHeap) Len() int {
	return len(iheap)
}

func (iheap integerHeap) Less(i, j int) bool {
	return iheap[i] < iheap[j]
}

func (iheap integerHeap) Swap(i, j int) {
	iheap[i], iheap[j] = iheap[j], iheap[i]
}

func (iheap *integerHeap) Push(i any) {
	*iheap = append(*iheap, i.(int))
}

func (iheap *integerHeap) Pop() any {
	previous := *iheap
	n := len(previous)
	x1 := previous[n-1]
	*iheap = previous[0 : n-1]
	return x1
}

func heaps() {
	intHeap := &integerHeap{1, 4, 5}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	fmt.Printf("minimum: %d\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
}
