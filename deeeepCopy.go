package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// DeepCopy ...
func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

// User ...
type User struct {
	Ids    []string
	Scores []float32
	ReqID  string
}

func dmain() {
	user1 := new(User)
	user1.Ids = []string{"a", "b", "c"}
	user1.Scores = []float32{1.2, 3.4, 4.2}
	user1.ReqID = "123dsfs"
	user2 := new(User)
	DeepCopy(user2, user1)
	user2.ReqID = "asssss123"
	fmt.Println(user1)
	fmt.Println(user2)

}
