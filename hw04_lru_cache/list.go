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

func NewListItem(value interface{}) *ListItem {
	return &ListItem{Value: value}
}

type list struct {
	len   int
	front *ListItem
	back  *ListItem
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
	listItem := NewListItem(v)
	defer incLen(l)

	if l.front == nil {
		l.front = listItem
		l.back = listItem

		return listItem
	}

	if l.len == 1 {
		l.back.Prev = listItem
	}

	listItem.Next = l.front
	l.front.Prev = listItem
	l.front = listItem

	return listItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	listItem := NewListItem(v)
	defer incLen(l)

	if l.front == nil {
		l.front = listItem
		l.back = listItem

		return listItem
	}

	listItem.Prev = l.back
	l.back.Next = listItem
	l.back = listItem

	return listItem
}

func (l *list) Remove(i *ListItem) {
	if i == l.front {
		l.front = i.Next
		i.Next.Prev = nil
	} else if i == l.back {
		l.back = i.Prev
		l.back.Next = nil

		i.Prev.Next = nil
		i.Prev = nil
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return new(list)
}

func incLen(l *list) {
	l.len++
}
