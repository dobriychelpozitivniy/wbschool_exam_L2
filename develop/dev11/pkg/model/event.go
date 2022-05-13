package model

type CreateEvent struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	DateAdded   string `json:"date_added" form:"date_added"`
	DateTodo    string `json:"date_todo" form:"date_todo"`
}

type Event struct {
	ID          string `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	DateAdded   string `json:"date_added" form:"date_added"`
	DateTodo    string `json:"date_todo" form:"date_todo"`
}
