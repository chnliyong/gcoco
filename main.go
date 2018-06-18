package main

import (
	"fmt"
	"github.com/gliderlabs/ssh"
	"io"
)

func main() {
	fmt.Println("Hello world")
	ssh.Handle(func(s ssh.Session) {
		io.WriteString(s, "Hello world\n")
	})
}
