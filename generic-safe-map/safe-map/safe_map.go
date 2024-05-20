package safemap

import (
	"fmt"
	"sync"
)

type SafeMap[K comparable, V any] struct {
	mutex sync.RWMutex
	data  map[K]V
}

func New[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		data: make(map[K]V),
	}
}

func (m *SafeMap[K, V]) Insert(key K, value V) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.data[key] = value
}

func (m *SafeMap[K, V]) Get(key K) (V, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	value, ok := m.data[key]
	if !ok {
		return value, fmt.Errorf("key: %v not found", key)
	}
	return value, nil
}

func (m *SafeMap[K, V]) Update(key K, value V) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, ok := m.data[key]; !ok {
		return fmt.Errorf("key: %v not found", key)
	}
	m.data[key] = value
	return nil
}

func (m *SafeMap[K, V]) Delete(key K) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, ok := m.data[key]; !ok {
		return fmt.Errorf("key: %v not found", key)
	}

	delete(m.data, key)
	return nil
}

func (m *SafeMap[K, V]) HasKey(key K) bool {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	_, ok := m.data[key]
	return ok
}
