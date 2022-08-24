package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Handler struct {
	Storage Storage
}

func NewHandler(storage Storage) *Handler {
	return &Handler{
		Storage: storage,
	}
}

func (h *Handler) createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		userEventDTO, errJson := parseBody(r.Body)
		if errJson != nil {
			b, _ := json.Marshal(NewInputDataError(errJson.Error()))
			w.WriteHeader(http.StatusBadRequest)
			w.Write(b)
		} else {
			user, err := h.Storage.CreateEvent(
				userEventDTO.Id,
				userEventDTO.Header,
				userEventDTO.Description,
				userEventDTO.Date,
			)
			if err != nil {
				b, _ := json.Marshal(NewInputDataError(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(b)
			} else {
				b, _ := json.Marshal(Result{User: *user})
				w.Write(b)
			}
		}
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		b, _ := json.Marshal(NewInputDataError(r.Method + " method not supported"))
		w.Write(b)
	}
}

func (h *Handler) updateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		userEventDTO, errJson := parseBody(r.Body)
		if errJson != nil {
			b, _ := json.Marshal(NewInputDataError(errJson.Error()))
			w.WriteHeader(http.StatusBadRequest)
			w.Write(b)
		} else {
			user, err := h.Storage.UpdateEvent(
				userEventDTO.Id,
				userEventDTO.Header,
				userEventDTO.Description,
				userEventDTO.Date,
			)
			if err != nil {
				b, _ := json.Marshal(NewInputDataError(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(b)
			} else {
				b, _ := json.Marshal(Result{User: *user})
				w.Write(b)
			}
		}
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		b, _ := json.Marshal(NewInputDataError(r.Method + " method not supported"))
		w.Write(b)
	}
}

func (h *Handler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		userEventDTO, errJson := parseBody(r.Body)
		if errJson != nil {
			b, _ := json.Marshal(NewInputDataError(errJson.Error()))
			w.WriteHeader(http.StatusBadRequest)
			w.Write(b)
		} else {
			user, err := h.Storage.DeleteEvent(
				userEventDTO.Id,
				userEventDTO.Date,
			)
			if err != nil {
				b, _ := json.Marshal(NewInputDataError(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(b)
			} else {
				b, _ := json.Marshal(Result{User: *user})
				w.Write(b)
			}
		}
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		b, _ := json.Marshal(NewInputDataError(r.Method + " method not supported"))
		w.Write(b)
	}
}

func (h *Handler) eventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		eventDto, errParse := parseBodyEventFor(r.URL.Query())
		if errParse != nil {
			b, _ := json.Marshal(NewInputDataError(errParse.Error()))
			w.WriteHeader(http.StatusBadRequest)
			w.Write(b)
		} else {
			var events []Event = h.Storage.EventsForDay(eventDto.Id, eventDto.Date)
			b, _ := json.Marshal(ResultEvent{Events: events})
			w.Write(b)
		}
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		b, _ := json.Marshal(NewInputDataError(r.Method + " method not supported"))
		w.Write(b)
	}
}

func (h *Handler) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		eventDto, errParse := parseBodyEventFor(r.URL.Query())
		if errParse != nil {
			b, _ := json.Marshal(NewInputDataError(errParse.Error()))
			w.WriteHeader(http.StatusBadRequest)
			w.Write(b)
		} else {
			var events []Event = h.Storage.EventsForWeek(eventDto.Id, eventDto.Date)
			b, _ := json.Marshal(ResultEvent{Events: events})
			w.Write(b)
		}
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		b, _ := json.Marshal(NewInputDataError(r.Method + " method not supported"))
		w.Write(b)
	}
}

func (h *Handler) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		eventDto, errParse := parseBodyEventFor(r.URL.Query())
		if errParse != nil {
			b, _ := json.Marshal(NewInputDataError(errParse.Error()))
			w.WriteHeader(http.StatusBadRequest)
			w.Write(b)
		} else {
			var events []Event = h.Storage.EventsForMonth(eventDto.Id, eventDto.Date)
			b, _ := json.Marshal(ResultEvent{Events: events})
			w.Write(b)
		}
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		b, _ := json.Marshal(NewInputDataError(r.Method + " method not supported"))
		w.Write(b)
	}
}

func parseBody(b io.ReadCloser) (*userEventDTO, error) {
	body, err := ioutil.ReadAll(b)
	defer b.Close()
	if err != nil {
		return nil, err
	}
	var ue userEventDTO
	errJson := json.Unmarshal(body, &ue)
	if errJson != nil {
		return nil, errJson
	}
	errValid := validateBody(&ue)
	return &ue, errValid
}

func validateBody(body *userEventDTO) error {
	if body.Id < 0 {
		return errors.New("id not valid")
	}
	if _, err := GetDate(body.Date); err != nil {
		return errors.New("not valid date")
	}
	return nil
}

func parseBodyEventFor(val url.Values) (*userEventForDTO, error) {
	id, err := strconv.Atoi(val.Get("id"))
	if err != nil {
		return nil, err
	}
	date, err1 := strconv.Atoi(val.Get("date"))
	if err1 != nil {
		return nil, err1
	}
	return &userEventForDTO{
		Id:   id,
		Date: date,
	}, nil
}

type userEventForDTO struct {
	Id   int `json:"id"`
	Date int `json:"date"`
}

type userEventDTO struct {
	Id          int    `json:"id"`
	Header      string `json:"header"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

type InputDataError struct {
	Err string `json:"error"`
}

func (i *InputDataError) Error() string {
	return i.Err
}

func NewInputDataError(msg string) *InputDataError {
	return &InputDataError{
		Err: msg,
	}
}

type Result struct {
	User User `json:"result"`
}

type ResultEvent struct {
	Events []Event `json:"result"`
}
