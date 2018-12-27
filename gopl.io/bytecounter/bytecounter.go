package main

import (
	"fmt"
)

//ByteCounter is a byte counter
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) //convert int to ByteCounter
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("Hello"))
	fmt.Println(c)
	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}
