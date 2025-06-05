# Casting []byte to string in golang

Most of the time doing `string(buffer)` where buffer is `[]byte` is fine but every time we do that we allocate new memory for the string and then copy everything from the buffer to the string. What if we want to avoid that?

Golang provides the unsafe package which we can use for casting pointers, so we can cast []byte (which has the pointer underneath) to a string but we have to know the memory layout of both of them.

### How and why this works

[]byte is a slice of bytes and all slices in go have 3 things: Pointer to the allocated memory, length, and capacity. Slices are mutable and we can use append, for example, to extend them. We could say that []byte looks something like this:

```go
type ByteSliceHeader struct {
    Ptr *byte
    Len uint
    Cap uint
}
```

This structure is commonly known as a slice header but instead of *byte it has *T for []T.

String header looks like this:

```go
type StringHeader struct {
    Ptr *byte
    Len uint
}
```

This means that if we have a pointer to a ByteSliceHeader we also have a pointer to a StringHeader because the first 2 fields match. If we were to get a pointer to ByteSliceHeader, cast it to StringHeader and then dereference it, we would get the first 2 fields copied to StringHeader. Keep in mind that StringHeader and ByteSliceHeader are allocated on the stack but the data underneath is allocated on the heap probably (doesn't have to be as you create []byte on the stack).

Because pointers are just memory addresses, when you dereference a pointer you tell your computer "okay from this address read X bytes". If you dereference an int32, X would be 4; if you dereference some BigStruct, the X would be sizeof(BigStruct). That's why, by the way, you cannot dereference a void* in C, because when you do it, C doesn't know how many bytes to read from that address.

So in our case, when we tell our computer to read StringHeader from some address, we are telling it to read just the 12 bytes (more realistically 16 bytes, because of the memory alignment and padding of the struct fields). The ByteSliceHeader is also 16 bytes wide so we are good. That's how you can cast []byte to string.

To be more precise, you can cast anything to anything, but how the memory will be interpreted and whether you, by doing that, try to read more bytes than are allocated and get a SEGFAULT is on you.

> Note: Casing back from string to []byte would just need from you to set the capacity to be the same as length.

## Examples

You can find examples in Go and C located in pointercast.go and pointercast.c respectively.

To run them you need to have gcc and go and make installed. 

- Run C example by typing `make c`
- Run GO example by typing `make go`
