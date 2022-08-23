package main

import (
	"errors"
	"sync"
	"time"
)

type Storage interface {
	CreateEvent(id int, header string, description string, date string) (*User, error)
	UpdateEvent(id int, header string, description string, date string) (*User, error)
	DeleteEvent(id int, date string) (*User, error)
	EventsForDay()
	EventsForWeek()
	EventsForMonth()
}

type MemoryStorage struct {
	sync.Mutex
	users map[int]User
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		users: make(map[int]User),
	}
}

func (m *MemoryStorage) CreateEvent(id int, header string, description string, date string) (*User, error) {
	m.Lock()
	if val, ok := m.users[id]; ok {
		return &val, val.CreateEvent(header, description, date)
	}
	var user User = User{
		Id:     id,
		Events: make(map[time.Time]Event),
	}
	err := user.CreateEvent(header, description, date)
	if err != nil {
		return nil, err
	}
	m.users[user.Id] = user
	m.Unlock()
	return &user, nil
}

func (m *MemoryStorage) UpdateEvent(id int, header string, description string, date string) (*User, error) {
	m.Lock()
	if _, ok := m.users[id]; !ok {
		return nil, errors.New("User does not exists")
	}
	var user User = m.users[id]
	m.Unlock()
	return &user, user.UpdateEvent(header, description, date)
}

func (m *MemoryStorage) DeleteEvent(id int, date string) (*User, error) {
	m.Lock()
	if _, ok := m.users[id]; !ok {
		return nil, errors.New("User does not exists")
	}
	var user User = m.users[id]
	m.Unlock()
	return &user, user.DeleteEvent(date)
}

func (m *MemoryStorage) EventsForDay()   {}
func (m *MemoryStorage) EventsForWeek()  {}
func (m *MemoryStorage) EventsForMonth() {}
