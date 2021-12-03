# Arrays and Slices

URL: [go.dev](https://go.dev/blog/slices)

## Arrays

Taking this declaration `var buffer [256]byte` we 
- declare a variable named buffer 
- set it's size/length to 256 bytes

We can check it's length by calling `len(buffer)` 


## Slices 

A slice is a data structure describing a contiguous (connecting) section of an array stored seperately from the slice variable itself.
A slice is not an array. A slice describes a piece of an array.


Given our buffer arrray we can create a slice 
```go
var slice []byte = buffer[100:150]
```

A more idiomatic way: 
```go
var slice = buffer[100:150]
```

Short declaration: 
```go
slice := buffer[100:150]
```

A slice has a length, and a pointer to an element of an array, like this

```go
type sliceHeader struct {
    Length  int
    ZerothElement *byte
}

slice := sliceHeader{
    Length: 50,
    ZerothElement: &buffer[100],
}
```

We can also slice a slice

```go
slice := buffer[100:150]
```

### Passing slices to functions

It's impotant to understand that a slice is itself a value, its a value holding a pointer and a length.

When we called IndexRune it was passed as a **copy of the slice header**.
``` go
func AddOneToEachElement(slice []byte) {
    for i := range slice {
        slice[i]++
    }
}
```

