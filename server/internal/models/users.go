package model

type User struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Id          int    `json:"id"`
	Rating      int    `json:"rating"`
	TasksSolved int    `json:"tasksCount"`
}
