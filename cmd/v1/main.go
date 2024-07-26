package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

// 这是一个点子，是未经测试的思路代码
// This is an idea, untested code.
func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run proxy.go <listen_addr> <target_ip> <target_port>")
		os.Exit(1)
	}

	listenAddr := os.Args[1]
	targetIP := os.Args[2]
	targetPort := os.Args[3]

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer ln.Close()
	fmt.Println("Listening on", listenAddr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}

		go handleConnection(conn, targetIP, targetPort)
	}
}

func handleConnection(local *net.TCPConn, targetIP string, targetPort string) {
	defer local.Close()

	targetAddr := fmt.Sprintf("%s:%s", targetIP, targetPort)
	target, err := net.Dial("tcp", targetAddr)
	if err != nil {
		fmt.Println("Error dialing:", err.Error())
		return
	}
	defer target.Close()

	go io.Copy(target, local)
	io.Copy(local, target)
}
