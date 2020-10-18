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
	l := new(ListItem)
	l.Value = 0
	l.Next = nil
	l.Prev = nil
	return l
}

func (l *ListItem) Len() int {
	return l.Value.(int)
}

func (l *ListItem) Front() *ListItem {
	return l.Prev
}

func (l *ListItem) Back() *ListItem {
	return l.Next
}

func (l *ListItem) PushFront(v interface{}) *ListItem {
	l.Prev = &ListItem{Value: v, Prev: nil, Next: l.Prev}
	if l.Len() == 0 {
		l.Next = l.Prev
	} else {
		l.Prev.Next.Prev = l.Prev
	}
	l.Value = l.Value.(int) + 1
	return l.Prev
}

func (l *ListItem) PushBack(v interface{}) *ListItem {
	l.Next = &ListItem{Value: v, Prev: l.Next, Next: nil}
	if l.Len() == 0 {
		l.Prev = l.Next
	} else {
		l.Next.Prev.Next = l.Next
	}
	l.Value = l.Value.(int) + 1
	return l.Next
}

func (l *ListItem) Remove(i *ListItem) {
	switch l.Len() {
	case 0: // Если пустой список, пропускаем
	case 1: // Если один элемент, делаем пустой список
		l.Prev, l.Next = nil, nil
		l.Value = 0
	case 2: // Если элементов два
		if l.Front() == i { // Если элемент первый
			l.Prev, l.Next = i.Next, i.Next
		} else { // Если элемент второй
			l.Prev, l.Next = i.Prev, i.Prev
		}
		l.Value = 1
		l.Prev.Prev, l.Next.Next = nil, nil // В оставшемся элементе удаляем ссылки на другие элементы
	default: // Когда элементов больше двух
		switch {
		case l.Front() == i: // Если элемент первый
			l.Prev = i.Next
			l.Prev.Prev = nil
		case l.Back() == i: // Если элемент последний
			l.Next = i.Prev
			l.Next.Next = nil
		default:
			i.Prev.Next, i.Next.Prev = i.Next, i.Prev
		}
		l.Value = l.Value.(int) - 1
	}
	i.Prev, i.Next = nil, nil // Удаляю ссылки на элементы в удаленном элементе
}

func (l *ListItem) MoveToFront(i *ListItem) {
	switch l.Len() {
	case 0 | 1: // Список пуст или один элемент
	case 2: // Два элемента и это не первый
		if i.Prev != nil {
			l.Prev, l.Next = l.Next, l.Prev
			l.Prev.Prev, l.Prev.Next = l.Prev.Next, l.Prev.Prev
			l.Next.Prev, l.Next.Next = l.Next.Next, l.Next.Prev
		}
	default:
		if l.Front() != i { // это не первый элемент
			tmpP := i.Prev
			if i.Next != nil { // Это не последний элемент
				tmpN := i.Next
				tmpP.Next, tmpN.Prev = tmpN, tmpP
			} else {
				tmpP.Next, l.Next = nil, tmpP
			}
			i.Prev, i.Next = nil, l.Front()
			i.Next.Prev, l.Prev = i, i
		}
	}
}
