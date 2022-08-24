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
	EventsForDay(id int, day int) []Event
	EventsForWeek(id int, week int) []Event
	EventsForMonth(id int, month int) []Event
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

func (m *MemoryStorage) EventsForDay(id int, day int) []Event {
	var user User = m.users[id]
	var days []Event = make([]Event, 0)
	for key, val := range user.Events {
		if key.Day() == day {
			days = append(days, val)
		}
	}
	return days
}
func (m *MemoryStorage) EventsForWeek(id int, week int) []Event {
	var user User = m.users[id]
	var weeks []Event = make([]Event, 0)
	for key, val := range user.Events {
		if int(key.Weekday()) == week {
			weeks = append(weeks, val)
		}
	}
	return weeks
}
func (m *MemoryStorage) EventsForMonth(id int, mnth int) []Event {
	var user User = m.users[id]
	var month []Event = make([]Event, 0)
	for key, val := range user.Events {
		if key.Month() == time.Month(mnth) {
			month = append(month, val)
		}
	}
	return month
}
