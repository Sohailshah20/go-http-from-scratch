package main

import (
	"fmt"
	"strings"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	port := "localhost:4221"
	// port := "0.0.0.0:4221"
	l, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	req := make([]byte, 1024)
	conn.Read(req)
	str := string(req)
	fmt.Println("req path ", str)
	clrf := strings.Split(str, "\r\n")
	first := strings.Split(clrf[0], " ")
	var path string
	execPath := strings.Split(first[1], "/")[1]
	if len(execPath) > 0 {
		path = execPath
	} else {
		path = "/"
	}
	if path == "/" {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
		conn.Close()
		return
	} else if path == "echo" {
		param := strings.Split(path, "/")[2]
		res := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(param), param)
		conn.Write([]byte(res))
		conn.Close()
		return
	} else if path == "user-agent" {
		clrf := strings.Split(str, "\r\n")
		headers := make(map[string]string)
		for i := 1; i < len(clrf)-2; i++ {
			split := strings.SplitN(clrf[i], ":", 2)
			headers[strings.TrimLeft(split[0], " ")] = strings.TrimLeft(split[1], " ")
		}
		res := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(headers["User-Agent"]), headers["User-Agent"])
		conn.Write([]byte(res))
		conn.Close()
		return

	} else {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		conn.Close()
		return
	}
}
