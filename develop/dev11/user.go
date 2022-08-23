package main

import (
	"errors"
	"time"
)

type User struct {
	Id     int
	Events map[time.Time]Event
}

func (u *User) CreateEvent(header string, description string, date string) error {
	var event Event = Event{}
	event.SetHeader(header)
	event.SetDescription(description)
	err := event.SetDate(date)
	if err != nil {
		return err
	}
	if _, ok := u.Events[event.Date]; ok {
		return errors.New("Event already exists")
	} else {
		u.Events[event.Date] = event
	}
	return nil
}

func (u *User) UpdateEvent(header string, description string, date string) error {
	t, err := GetDate(date)
	if err != nil {
		return err
	}
	if val, ok := u.Events[t]; ok {
		val.SetHeader(header)
		val.SetDescription(description)
		val.SetDate(date)
	} else {
		return errors.New("Event does not exists")
	}
	return nil
}

func (u *User) DeleteEvent(date string) error {
	t, err := GetDate(date)
	if err != nil {
		return errors.New("Event does not exists")
	}
	delete(u.Events, t)
	return nil
}
