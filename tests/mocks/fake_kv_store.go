package mocks

import "errors"

// Fake implementation of the KVStore interface
type FakeKVStore struct {
	Data map[string]string
}

func NewFakeKVStore() *FakeKVStore {
	return &FakeKVStore{
		Data: make(map[string]string),
	}
}

func (m *FakeKVStore) Set(key string, value string) (string, error) {
	m.Data[key] = value
	return value, nil
}

func (m *FakeKVStore) Get(key string) (string, error) {
	value, exists := m.Data[key]
	if !exists {
		return "", errors.New("Key not found")
	}
	return value, nil
}

func (m *FakeKVStore) Delete(key string) (int, error) {
	if _, exists := m.Data[key]; exists {
		delete(m.Data, key)
		return 1, nil
	}
	return 0, errors.New("Key not found")
}
