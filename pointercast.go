package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"unsafe"
)

type User struct {
	Username string
	Email    string
}

func main() {
	cExmaple()
}

func cExmaple() {
	buffer := make([]byte, 0, 10)
	buffer = append(buffer, 'A')
	buffer = append(buffer, 'B')
	buffer = append(buffer, 'C')
	buffer = append(buffer, 0)

	s := *(*string)(unsafe.Pointer(&buffer))
	fmt.Printf("\"%s\" %d\n", s, len(s))

	buffer[2] = 'D'
	fmt.Printf("\"%s\" %d\n", s, len(s))
}

func jsonExample() {
	users := []User{
		{Username: "Dusan", Email: "dukidjordjic@gmail.com"},
		{Username: "Cone", Email: "conedjordjic@gmail.com"},
	}

	data, _ := json.Marshal(users)
	s := (*string)(unsafe.Pointer(&data))
	fmt.Printf("\"%s\" %d %d %d\n", *s, len(*s), len(data), cap(data))

	data = []byte{65, 66, 67, 49, 50, 51}
	s = (*string)(unsafe.Pointer(&data))
	fmt.Printf("\"%s\" %d %d %d\n", *s, len(*s), len(data), cap(data))

	data = nil
	s = (*string)(unsafe.Pointer(&data))
	fmt.Printf("\"%s\" %d %d %d\n", *s, len(*s), len(data), cap(data))

	b := strings.Builder{}
	json.NewEncoder(&b).Encode(users)
	ss := b.String()

	fmt.Printf("\"%s\" %d %d %d\n", ss, len(ss), len(data), cap(data))
}
