package models

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Status   int    `json:"status"` //0 - inactive , 1- active
}
