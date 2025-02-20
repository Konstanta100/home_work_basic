package entity

import "encoding/json"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u User) String() string {
	body, _ := json.Marshal(u)
	return string(body)
}
