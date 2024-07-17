package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "aaaa \r\nbb\r\ncc\r\n\r\n"
	fmt.Println(strings.Split(str, "\r\n"))

}
