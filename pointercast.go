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
	fmt.Println("===== C Example ===========")
	cExample()
	fmt.Println("===========================")
	fmt.Println("===== Safer Example =======")
	saferExample()
	fmt.Println("===========================")
	fmt.Println("===== JSON Example ========")
	jsonExample()
	fmt.Println("===========================")
}

// This is a safer way of doing it, because they can change how strings or slices are
// represented underneath as its mentioned in reflect/value.go
func saferExample() {
	{
		buffer := make([]byte, 0, 10)
		buffer = append(buffer, 'A', 'B', 'C', 0)
		// strings in go do not need 0 at the end but we are keeping it to
		// have the same output as C code

		l := len(buffer)
		bytes := unsafe.SliceData(buffer)
		s := unsafe.String(bytes, l)

		fmt.Printf("\"%s\" %d\n", s, len(s))

		buffer[2] = 'D'
		fmt.Printf("\"%s\" %d\n", s, len(s))

	}

	{
		buffer := []byte{65, 66, 67}

		l := len(buffer)
		bytes := unsafe.SliceData(buffer)
		s := unsafe.String(bytes, l)

		fmt.Printf("\"%s\" %d\n", s, len(s))
	}

	{
		var buffer []byte = nil

		l := len(buffer)
		bytes := unsafe.SliceData(buffer)
		s := unsafe.String(bytes, l)

		fmt.Printf("\"%s\" %d\n", s, len(s))
	}
}

func cExample() {
	buffer := make([]byte, 0, 10)
	buffer = append(buffer, 'A', 'B', 'C', 0)
	// strings in go do not need 0 at the end but we are keeping it to
	// have the same output as C code

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

	// Its possible to use strings.Builder beacuse it implements Write
	// and the data is casted to string underneath using unsafe.
	//
	// func (b *Builder) String() string {
	// 	return unsafe.String(unsafe.SliceData(b.buf), len(b.buf))
	// }

	b := strings.Builder{}
	json.NewEncoder(&b).Encode(users)
	ss := b.String()

	fmt.Printf("\"%s\" %d %d %d\n", ss, len(ss), len(data), cap(data))
}
