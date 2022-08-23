package main

import "time"

type Event struct {
	Date        time.Time
	Header      string
	Description string
}

func (e *Event) SetHeader(head string) {
	e.Header = head
}

func (e *Event) SetDescription(desc string) {
	e.Description = desc
}

func (e *Event) SetDate(date string) error {
	t, err := GetDate(date)
	if err != nil {
		return err
	}
	e.Date = t
	return nil
}
