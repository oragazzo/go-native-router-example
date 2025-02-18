package models

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserList struct {
	Users []User `json:"users"`
}
