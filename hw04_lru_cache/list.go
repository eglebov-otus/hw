package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int
	Front() *listItem
	Back() *listItem
	PushFront(v interface{}) *listItem
	PushBack(v interface{}) *listItem
	Remove(i *listItem)
	MoveToFront(i *listItem)
}

type listItem struct {
	Prev, Next *listItem
	Value      interface{}
}

type list struct {
	len        int
	head, tail *listItem
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *listItem {
	return l.head
}

func (l list) Back() *listItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *listItem {
	newItem := &listItem{Value: v}
	curHead := l.head

	if curHead != nil {
		newItem.Prev = curHead
		curHead.Next = newItem
		l.head = newItem
	} else {
		l.head = newItem
		l.tail = newItem
	}

	l.len++

	return newItem
}

func (l *list) PushBack(v interface{}) *listItem {
	newItem := &listItem{Value: v}
	curTail := l.tail

	if curTail != nil {
		newItem.Next = curTail
		curTail.Prev = newItem
		l.tail = newItem
	} else {
		l.head = newItem
		l.tail = newItem
	}

	l.len++

	return newItem
}

func (l *list) Remove(i *listItem) {
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.head = i.Prev
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.tail = i.Next
	}

	l.len--
}

func (l *list) MoveToFront(i *listItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return &list{}
}
