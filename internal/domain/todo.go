package domain

type Todo struct {
	ID        int `json:"id"`
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
}
