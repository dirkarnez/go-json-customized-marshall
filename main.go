package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type UUIDEx uuid.UUID

// MarshalJSON marshals the UUIDEx type to a JSON UUID string.
func (my UUIDEx) MarshalJSON() ([]byte, error) {
	u := uuid.UUID(my)
	return []byte(fmt.Sprintf(`"%s"`, strings.ReplaceAll(u.String(), "-", ""))), nil
}

func NewUUIDEx() UUIDEx {
	return UUIDEx(uuid.New())
}

type User struct {
	ID   UUIDEx `json:"id"`
	Name string `json:"name"`
}

func (f *User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		Name string `json:"name_extra"`
		*Alias
	}{
		Name:  f.Name + "!",
		Alias: (*Alias)(f),
	})
}

func main() {
	user := &User{ID: NewUUIDEx(), Name: "Frank"}
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println(string(b))
	/*
		{"name_extra":"Frank!","id":"e9af4376128a43eb9c72ce0a08886594","name":"Frank"}
	*/
}
