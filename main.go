package main

import (
	"github.com/google/uuid"
	"encoding/json"
	"fmt"
	"strings"
)

type UUIDEx uuid.UUID

// MarshalJSON marshals the UUIDEx type to a JSON UUID string.
func (my UUIDEx) MarshalJSON() ([]byte, error) {
	u := uuid.UUID(my)
	return []byte(fmt.Sprintf(`"%s"`, strings.ReplaceAll(u.String(), "-", ""))), nil
}


type User struct {
	ID UUIDEx
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
