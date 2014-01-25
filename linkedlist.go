package datastructures

import "fmt"

type LinkedList struct {
	head   *ListItem
	length int
}

type ListItem struct {
	item interface{}
	next *ListItem
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil}
}

func (l *LinkedList) Insert(value interface{}) {
	item := &ListItem{item: value, next: l.head}
	l.length++
	l.head = item
}

func (l *LinkedList) Length() int {
	return l.length
}

func (l *LinkedList) PrintAll() {

	value := l.head
	for value != nil {
		fmt.Println(value.item)
		value = value.next

	}
}

func (l *LinkedList) Traverse(f func(item *ListItem)) {

	value := l.head
	for value != nil {
		if f != nil {
			f(value)
		}

		value = value.next
	}
}

func (l *LinkedList) Search(query interface{}) bool {
	if query == "" {
		return false
	}

	value := l.head
	for value != nil {
		if value.item == query {
			return true
		}
		value = value.next
	}
	return false
}

func (l *LinkedList) DeleteItem(query interface{}) {
	var previous *ListItem
	value := l.head
	for value != nil {
		if value.item == query {
			if previous != nil {
				previous.next = value.next
			} else {
				l.head = value.next
			}

		}
		previous = value
		value = value.next
	}

}

/*func main() {
	list := linkedList{head: nil}
	list.insert("hello")
	list.insert("world")
	list.printAll()

	list.deleteItem("hello")
	list.printAll()
}
*/
