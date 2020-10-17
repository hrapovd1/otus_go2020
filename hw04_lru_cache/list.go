package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int                          // длина списка
	Front() *ListItem                  // первый элемент списка
	Back() *ListItem                   // последний элемент списка
	PushFront(v interface{}) *ListItem // добавить значение в начало
	PushBack(v interface{}) *ListItem  // добавить значение в конец
	Remove(i *ListItem)                // удалить элемент
	MoveToFront(i *ListItem)           // переместить элемент в начало
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

func NewList() *ListItem {
	return &ListItem{Value: 0, Prev: nil, Next: nil}
}

func (l ListItem) Len() int {
	return l.Value.(int)
}

func (l ListItem) Front() *ListItem {
	return l.Prev
}

func (l ListItem) Back() *ListItem {
	return l.Next
}

func (l ListItem) PushFront(v interface{}) *ListItem {
	l.Prev = &ListItem{Value: v, Prev: nil, Next: l.Prev}
	if l.Len() == 0 {
		l.Next = l.Prev
	}
	l.Value = l.Value.(int) + 1
	return l.Prev
}

func (l ListItem) PushBack(v interface{}) *ListItem {
	l.Next = &ListItem{Value: v, Next: nil, Prev: l.Next}
	if l.Len() == 0 {
		l.Prev = l.Next
	}
	l.Value = l.Value.(int) + 1
	return l.Next
}

func (l ListItem) Remove(i *ListItem) {
	switch l.Len() {
	case 0: // Если пустой список, пропускаем
	case 1: // Если один элемент, делаем пустой список
		l.Prev, l.Next = nil, nil
		l.Value = 0
	case 2: // Если элементов два
		if i.Prev == nil { // Если элемент первый
			l.Prev, l.Next = i.Next, i.Next
		} else if i.Next == nil { // Если элемент второй
			l.Prev, l.Next = i.Prev, i.Prev
		}
		l.Prev.Prev, l.Next.Next = nil, nil // В оставшемся элементе удаляем ссылки на другие элементы
		l.Value = 1
	default: // Когда элементов больше двух
		switch {
		case i.Prev == nil: // Если элемент первый
			i.Next.Prev = nil
		case i.Next == nil: // Если элемент последний
			i.Prev.Next = nil
		default:
			i.Prev.Next, i.Next.Prev = i.Next, i.Prev
		}
		l.Value = l.Value.(int) - 1
	}
}

func (l ListItem) MoveToFront(i *ListItem) {
	switch l.Len() {
	case 0 | 1: // Список пуст или один элемент
	case 2: // Два элемента и это не первый
		if i.Prev != nil {
			l.Prev, l.Next = l.Next, l.Prev
			l.Prev.Prev, l.Prev.Next = l.Prev.Next, l.Prev.Prev
			l.Next.Prev, l.Next.Next = l.Next.Next, l.Next.Prev
		}
	default:
		if l.Prev != i { // это не первый элемент
			tmpP := i.Prev
			if i.Next != nil { // Это не последний элемент
				tmpN := i.Next
				tmpP.Next, tmpN.Prev = tmpN, tmpP
			} else {
				tmpP.Next, l.Next = nil, tmpP
			}
			i.Prev, i.Next, l.Prev = nil, l.Prev, i
		}
	}
}
