package main

import (
	"testing"
	"time"
)

func TestUserCreateEvent(t *testing.T) {
	var header string = "head"
	var description string = "desc"
	var date string = "10-11-2022"
	var user User = User{
		Id:     1,
		Events: make(map[time.Time]Event),
	}
	err := user.CreateEvent(header, description, date)
	if err != nil {
		t.Error(err)
	}
	tm, _ := GetDate(date)
	if _, ok := user.Events[tm]; !ok {
		t.Error("key not exist")
	}
}

func TestUserCreateEventAgain(t *testing.T) {
	var header string = "head"
	var description string = "desc"
	var date string = "10-11-2022"
	var user User = User{
		Id:     1,
		Events: make(map[time.Time]Event),
	}
	user.CreateEvent(header, description, date)
	err := user.CreateEvent(header, description, date)
	if err != nil {
		t.Log(err)
	} else {
		t.Error("Same event created")
	}
}

func TestUserUpdateEvent(t *testing.T) {
	var header string = "head"
	var description string = "desc"
	var date string = "10-11-2022"
	var user User = User{
		Id:     1,
		Events: make(map[time.Time]Event),
	}
	err := user.CreateEvent(header, description, date)
	if err != nil {
		t.Error(err)
	}

	var headerCh string = "head1"
	var descriptionCh string = "desc1"
	err1 := user.UpdateEvent(headerCh, descriptionCh, date)
	if err1 != nil {
		t.Error(err1)
	}
}

func TestUserDeleteEvent(t *testing.T) {
	var header string = "head"
	var description string = "desc"
	var date string = "10-11-2022"
	var user User = User{
		Id:     1,
		Events: make(map[time.Time]Event),
	}
	err := user.CreateEvent(header, description, date)
	if err != nil {
		t.Error(err)
	}
	err1 := user.DeleteEvent(date)
	if err1 != nil {
		t.Error(err1)
	}
}
