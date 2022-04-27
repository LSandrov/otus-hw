package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len         int
	front, back *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	listItem := &ListItem{Value: v, Next: l.front}

	if l.front == nil {
		l.back = listItem
	} else {
		l.front.Prev = listItem
	}

	l.front = listItem
	l.len++

	return listItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	listItem := &ListItem{Value: v, Prev: l.back}

	if l.back == nil {
		l.back = listItem
	} else {
		l.back.Next = listItem
	}

	l.back = listItem
	l.len++

	return listItem
}

func (l *list) Remove(i *ListItem) {
	l.pos(i)

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.front == i {
		return
	}
	if l.back == i {
		l.back = i.Prev
		l.back.Next = nil
	} else {
		l.pos(i)
	}

	currentFront := l.front

	l.front = i
	l.front.Prev = nil
	l.front.Next = currentFront
	l.front.Next.Prev = i
}

func NewList() List {
	return new(list)
}

func (l *list) pos(i *ListItem) {
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.back = i.Prev
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.front = i.Next
	}
}
