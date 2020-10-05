package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int                          // длина списка
	Front() *listItem                  // первый элемент списка
	Back() *listItem                   // последний элемент списка
	PushFront(v interface{}) *listItem // добавить значение в начало
	PushBack(v interface{}) *listItem  // добавить значение в конец
	Remove(i *listItem)                // удалить элемент
	MoveToFront(i *listItem)           // переместить элемент в начало
}

type listItem struct {
	element int
	next    *listItem
	prev    *listItem
}

type list struct {
	list []listItem
}

/*
type list struct {
	current *listItem
	next    *listItem
	prev    *listItem
	count   *int
}
*/

func NewList() List {
	return list{make([]listItem, 0)}
}

func (l list) Len() int {
	return len(l.list)
}

func (l list) Front() *listItem {
	return &l.list[0]
}

func (l list) Back() *listItem {
	return &l.list[l.Len()-1]
}

func (l list) PushFront(v interface{}) *listItem {
	item := listItem{element: v.(int), prev: nil, next: &l.list[0]}
	l.list[0].prev = &item
	tmp := make([]listItem, l.Len()+1)
	tmp = append(tmp, item)
	l.list = append(tmp, l.list...)
	return &item
}

func (l list) PushBack(v interface{}) *listItem {
	item := listItem{element: v.(int), next: nil, prev: &l.list[l.Len()-1]}
	l.list[l.Len()-1].next = &item
	tmp := make([]listItem, l.Len()+1)
	tmp = append(tmp, l.list...)
	l.list = append(tmp, item)
	return &item
}

func (l list) Remove(i *listItem) {
	//TODO: need Remove code
	if l.Len() != 0 {
		tmp := make([]listItem, l.Len()-1)

	}
}

func (l list) MoveToFront(i *listItem) {
	//TODO: need MoveToFront code
	switch l.Len() {
	case 1:
	case 2:
	}
}
