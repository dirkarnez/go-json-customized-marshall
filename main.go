package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
}

func (f *User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		Name string
		*Alias
	}{
		Name:  f.Name + "!",
		Alias: (*Alias)(f),
	})
}
func main() {
	user := &User{Name: "Frank"}
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println(string(b))
}
