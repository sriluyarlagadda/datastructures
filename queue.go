package datastructures

type Queue struct {
	items []interface{}
	size  int
	front int
	back  int
}

func NewQueue(queueSize int) Queue {
	return Queue{items: make([]interface{}, queueSize), size: queueSize, front: 0, back: 0}
}

func (q *Queue) Enque(x interface{}) {
	if (q.front - q.back) < q.size {
		q.items[q.back] = x
		q.back++
	}
}

func (q *Queue) Deque() interface{} {
	if q.front == q.back {
		return nil
	}
	item := q.items[q.front]
	q.front++
	return item
}

func (q *Queue) Length() int {
	return (q.front - q.back)
}
