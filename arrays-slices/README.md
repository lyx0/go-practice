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
    var slice []byte = buffer[100:150]

A more idiomatic way: 
    `var slice = buffer[100:150]`.

Short declaration: 
    `slice := buffer[100:150]`

A slice has a length, and a pointer to an element of an array, like this

```
type sliceHeader struct {
    Length  int
    ZerothElement *byte
}

slice := sliceHeader{
    Length: 50,
    ZerothElement: &buffer[100],
}
```
