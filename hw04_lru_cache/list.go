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
}

type list struct {
	current *listItem
	next    *listItem
	prev    *listItem
	count   *int
}

func NewList() List {
	count := 0
	return &list{next: nil, prev: nil, current: &listItem{}, count: &count}
}

func (l list) Len() int {
	return *l.count
}

func (l list) Front() *listItem {
	return &listItem{}
}

func (l list) Back() *listItem {
	return &listItem{}
}

func (l list) PushFront(v interface{}) *listItem {
	newItem := listItem{element: v.(int)}
	*l.count++
	l.next = l.prev
	l.prev = nil
	return &newItem
}

func (l list) PushBack(v interface{}) *listItem {
	return &listItem{}
}

func (l list) Remove(i *listItem) {
}

func (l list) MoveToFront(i *listItem) {
}
