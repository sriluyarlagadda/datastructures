package datastructures

import (
	"fmt"
)

type Item struct {
	key   int
	value int
}

type PriorityQueue struct {
	items        []Item
	size         int
	currentCount int
}

type QueuFullError struct {
	size int
}

func (e *QueuFullError) Error() string {
	return fmt.Sprintf("Error:Queue size %d full!!!", e.size)
}

func NewPriorityQueue(size int) PriorityQueue {
	items := make([]Item, size)
	queueSize := size
	return PriorityQueue{items: items, size: queueSize, currentCount: 0}
}

func (q *PriorityQueue) Size() int {
	return q.length()
}

func (q *PriorityQueue) Capacity() int {
	return q.size
}

func (q *PriorityQueue) parent(n int) int {
	if n == 1 {
		return -1
	}
	return int(n / 2)
}

func (q *PriorityQueue) child(n int) int {
	return 2 * n
}

func (q *PriorityQueue) length() int {
	return q.currentCount
}

func (q *PriorityQueue) Insert(key int) error {

	item := Item{key: key, value: key}
	if q.length() >= q.size {
		return &QueuFullError{size: q.size}
	} else {
		count := q.length()
		count = count + 1
		q.currentCount++
		q.items[count] = item
		fmt.Println("insertion location:", count)
		fmt.Println("queue length:", q.length())
		bubbleUp(q, count)
	}
	return nil

}

func bubbleUp(q *PriorityQueue, itemPosition int) {
	if q.parent(itemPosition) == -1 {
		return
	}
	parentPosition := q.parent(itemPosition)
	if q.items[parentPosition].key > q.items[itemPosition].key {
		swap(q, itemPosition, parentPosition)
		bubbleUp(q, parentPosition)
	}
}

func swap(q *PriorityQueue, firstPosition int, secondPosition int) {
	q.items[firstPosition], q.items[secondPosition] = q.items[secondPosition], q.items[firstPosition]
}

func (q *PriorityQueue) GetMinimum() int {
	min := -1
	if q.length() == 0 {
		fmt.Println("emtpy queue")
		return min
	}
	min = q.items[1].key
	q.items[1] = q.items[q.length()]
	q.currentCount--
	bubbleDown(q, 1)
	return min
}

func bubbleDown(q *PriorityQueue, itemPosition int) {
	childPosition := q.child(itemPosition)
	var minChild int
	minChild = itemPosition

	for i := 1; i <= 2; i++ {

		if childPosition*i <= q.currentCount && q.items[minChild].key > q.items[childPosition*i].key {
			minChild = childPosition * i
		}
	}

	if minChild != itemPosition {
		swap(q, minChild, itemPosition)
		bubbleDown(q, minChild)
	}
}
