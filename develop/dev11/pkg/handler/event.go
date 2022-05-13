package handler

import (
	"dev11/pkg/model"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// POST /create_event
// POST /update_event
// POST /delete_event
// GET /events_for_day
// GET /events_for_week
// GET /events_for_month

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.Router.RespondWithError(w, 400, "not valid body")

		return
	}

	f, err := deserializeCreateEventPOST(r.PostForm)
	if err != nil {
		h.Router.RespondWithError(w, 400, err.Error())

		return
	}

	id, err := h.s.Create(*f)
	if err != nil {
		h.Router.RespondWithError(w, 503, err.Error())

		return
	}

	h.Router.RespondWithJSON(w, 200, map[string]interface{}{"result": id})
}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.Router.RespondWithError(w, 400, "not valid body")

		return
	}

	f, err := deserializeEventPOST(r.PostForm)
	if err != nil {
		h.Router.RespondWithError(w, 400, err.Error())

		return
	}

	e, err := h.s.Update(*f)
	if err != nil {
		h.Router.RespondWithError(w, 503, err.Error())

		return
	}

	h.RespondWithJSON(w, 200, map[string]interface{}{"result": e})
}

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.Router.RespondWithError(w, 400, "not valid body")

		return
	}

	id := r.Form.Get("id")
	if id == "" {
		h.Router.RespondWithError(w, 400, "empty body")

		return
	}

	id, err = h.s.Delete(id)
	if err != nil {
		h.Router.RespondWithError(w, 503, err.Error())

		return
	}

	h.RespondWithJSON(w, 200, map[string]interface{}{"result": id})
}

func (h *Handler) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	date, ok := r.URL.Query()["date"]
	if !ok || len(date[0]) < 1 {
		h.Router.RespondWithError(w, 400, "not valid body")

		return
	}

	if ok {
		err := validateDate(date[0])
		if err != nil {
			h.Router.RespondWithError(w, 400, err.Error())

			return
		}
	}

	e, err := h.s.GetForDay(date[0])
	if err != nil {
		h.Router.RespondWithError(w, 400, err.Error())

		return
	}

	h.RespondWithJSON(w, 200, map[string]interface{}{"result": e})
}

func (h *Handler) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	numberWeek, ok := r.URL.Query()["week"]
	if !ok || len(numberWeek[0]) < 1 {
		h.Router.RespondWithError(w, 400, "not valid body")

		return
	}

	nw, err := strconv.Atoi(numberWeek[0])
	if err != nil {
		h.Router.RespondWithError(w, 400, err.Error())

		return
	}

	if ok {
		if nw > 53 || nw < 1 {
			h.Router.RespondWithError(w, 400, "not valid body")

			return
		}
	}

	year, ok := r.URL.Query()["year"]
	if !ok || len(year[0]) < 1 {
		h.Router.RespondWithError(w, 400, fmt.Sprintf("Empty year"))

		return
	}

	y, err := strconv.Atoi(year[0])
	if err != nil {
		h.Router.RespondWithError(w, 400, err.Error())

		return
	}

	es, err := h.s.GetForWeek(nw, y)
	if err != nil {
		h.RespondWithError(w, http.StatusServiceUnavailable, err.Error())

		return
	}

	h.RespondWithJSON(w, 200, map[string]interface{}{"result": es})
}

func (h *Handler) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	numberMonth, ok := r.URL.Query()["month"]
	if !ok || len(numberMonth[0]) < 1 {
		h.Router.RespondWithError(w, 400, "not valid body")

		return
	}

	nm, err := strconv.Atoi(numberMonth[0])
	if err != nil {
		h.Router.RespondWithError(w, 400, err.Error())

		return
	}

	if ok {
		if nm > 53 || nm < 1 {
			h.Router.RespondWithError(w, 400, "not valid body")

			return
		}
	}

	year, ok := r.URL.Query()["year"]
	if !ok || len(year[0]) < 1 {
		h.Router.RespondWithError(w, 400, fmt.Sprintf("Empty year"))

		return
	}

	y, err := strconv.Atoi(year[0])
	if err != nil {
		h.Router.RespondWithError(w, 400, err.Error())

		return
	}

	es, err := h.s.GetForMonth(nm, y)
	if err != nil {
		h.RespondWithError(w, http.StatusServiceUnavailable, err.Error())

		return
	}

	h.RespondWithJSON(w, 200, map[string]interface{}{"result": es})
}

func deserializeCreateEventPOST(postForm url.Values) (*model.CreateEvent, error) {
	fields := make(map[string]string)

	name := postForm.Get("name")
	fields["name"] = name
	desc := postForm.Get("description")
	fields["description"] = desc
	da := postForm.Get("date_added")
	fields["date_added"] = da
	dt := postForm.Get("date_todo")
	fields["date_todo"] = dt

	for i, v := range fields {
		if strings.TrimSpace(v) == "" {
			return nil, fmt.Errorf("Empty %s field", i)
		}
	}

	err := validateDate(dt)
	if err != nil {
		return nil, fmt.Errorf("Incorrect date in date_todo: %s", err)
	}

	err = validateDate(da)
	if err != nil {
		return nil, fmt.Errorf("Incorrect date in date_added: %s", err)
	}

	return &model.CreateEvent{
		Name:        postForm.Get("name"),
		Description: postForm.Get("description"),
		DateAdded:   postForm.Get("date_added"),
		DateTodo:    postForm.Get("date_todo"),
	}, nil
}

func validateDate(date string) error {
	d := strings.Split(date, "-")
	if len(d) != 3 {
		return fmt.Errorf("day or month or year not exist")
	}

	for i := 0; i < 3; i++ {
		v, err := strconv.Atoi(d[i])
		if err != nil {
			return fmt.Errorf("Symbols in date")
		}

		if i == 0 {
			if (v < 0) || (v > 31) {
				return fmt.Errorf("Incorrect day")
			}
		}

		if i == 1 {
			if (v < 0) || (v > 12) {
				return fmt.Errorf("Incorrect month")
			}
		}

		if i == 2 {
			if v < 0 {
				return fmt.Errorf("Incorrect year")
			}
		}

	}

	return nil
}

func deserializeEventPOST(postForm url.Values) (*model.Event, error) {
	fields := make(map[string]string)

	id := postForm.Get("id")
	fields["id"] = id
	name := postForm.Get("name")
	fields["name"] = name
	desc := postForm.Get("description")
	fields["description"] = desc
	da := postForm.Get("date_added")
	fields["date_added"] = da
	dt := postForm.Get("date_todo")
	fields["date_todo"] = dt

	for i, v := range fields {
		if strings.TrimSpace(v) == "" {
			return nil, fmt.Errorf("Empty %s field", i)
		}
	}

	err := validateDate(dt)
	if err != nil {
		return nil, fmt.Errorf("Incorrect date in date_todo: %s", err)
	}

	err = validateDate(da)
	if err != nil {
		return nil, fmt.Errorf("Incorrect date in date_added: %s", err)
	}

	return &model.Event{
		ID:          id,
		Name:        postForm.Get("name"),
		Description: postForm.Get("description"),
		DateAdded:   postForm.Get("date_added"),
		DateTodo:    postForm.Get("date_todo"),
	}, nil
}
