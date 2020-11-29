package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type Cache interface {
	Set(key string, value interface{}) bool
	Get(key string) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	Key   Key
	Value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (lru *lruCache) Set(key string, value interface{}) bool {
	if item, exist := lru.items[Key(key)]; exist {
		lru.queue.MoveToFront(item)
		item.Value = cacheItem{Key: Key(key), Value: value}
		return true // Элемент уже присутствовал
	}

	if lru.queue.Len() == lru.capacity {
		// Если кеш полный, удаляю последний элемент
		lru.Clear()
	}

	itemTmp := cacheItem{Key: Key(key), Value: value}
	lru.items[Key(key)] = lru.queue.PushFront(itemTmp)

	return false // Элемент был вновь добавлен
}

func (lru *lruCache) Get(key string) (interface{}, bool) {
	if item, exist := lru.items[Key(key)]; exist {
		lru.queue.MoveToFront(item)
		return item.Value.(cacheItem).Value, true // Элемент присутствовал
	}
	return nil, false
}

func (lru *lruCache) Clear() {
	lastItem := lru.queue.Back()
	delete(lru.items, lastItem.Value.(cacheItem).Key)
	lru.queue.Remove(lastItem)
}
