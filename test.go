package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "GET /echo/hello HTTP/1.1\r\nUser-Agent: Mozilla/5.0 (Windows NT; Windows NT 10.0; en-US) WindowsPowerShell/5.1.22621.3880\r\nHost: localhost:4221\r\n\r\n"
	clrf := strings.Split(str, "\r\n")
	first := strings.Split(clrf[0], " ")
	// fmt.Println("params ", strings.Split(first[1], "/"))
	method := first[0]
	var path string
	execPath := strings.Split(first[1], "/")[1]
	if len(execPath) > 0 {
		path = execPath
	} else {
		path = "/"
	}
	headers := make(map[string]string)
	for i := 1; i < len(clrf)-2; i++ {
		split := strings.SplitN(clrf[i], ":", 2)
		headers[strings.TrimLeft(split[0], " ")] = strings.TrimLeft(split[1], " ")
	}
	req := req{
		method:  method,
		path:    path,
		headers: headers,
	}
	req.getPathParams(first)
	fmt.Printf("%+v", req)
}

type req struct {
	method  string
	path    string
	headers map[string]string
}

func (r req) getPathParams(path []string){
	fmt.Println("given path :", strings.Split(path[1], "/"))
}
