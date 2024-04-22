package main

import (
	"fmt"
	"os"
)

var server_chan_receive chan string

func main() {
	go writeFile()
}

func writeFile() {
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0o644)
	if err != nil {
		fmt.Println(err.Error())
	}
}
