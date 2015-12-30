package models

// My ToDo Entry structure
type Todo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}
