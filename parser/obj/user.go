package obj

import "fmt"

type User struct {
	Name    string         `json:"name"`
	Age     int            `json:"age"`
	Friends []string       `json:"friends"`
	Grade   map[string]int `json:"grade"`
	Brother Family         `json:"brother"`
}

type Family struct {
	Relation string `json:"relation"`
	Age      int    `json:"age"`
}

func (user *User) print()  {
	fmt.Println(user.Name)
}