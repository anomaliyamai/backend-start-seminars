package storage

import (
	"errors"
	"sync"
)

type RaiStorage struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewRaiStorage() *RaiStorage {
	return &RaiStorage{
		data: make(map[string]string),
	}
}

func (rs *RaiStorage) Get(key string) (*string, error) {
	rs.mu.RLock()
	defer rs.mu.RUnlock()

	value, exists := rs.data[key]
	if !exists {
		return nil, errors.New("key not found")
	}
	return &value, nil
}

func (rs *RaiStorage) Put(key string, value string) error {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	rs.data[key] = value
	return nil
}

func (rs *RaiStorage) Post(key string, value string) error {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	if _, exists := rs.data[key]; exists {
		return errors.New("key already exists")
	}
	rs.data[key] = value
	return nil
}

func (rs *RaiStorage) Delete(key string) error {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	if _, exists := rs.data[key]; !exists {
		return errors.New("key not found")
	}
	delete(rs.data, key)
	return nil
}
