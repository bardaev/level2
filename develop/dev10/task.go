package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	go StartServer()
	var timeout int = *flag.Int("timeout", 10, "Timeout in seconds")
	flag.Parse()

	var host string = flag.Arg(0)
	var port string = flag.Arg(1)

	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer conn.Close()

	sc := bufio.NewScanner(os.Stdin)
	// httpRequest := "GET /"
	for sc.Scan() {
		var req string = sc.Text()
		conn.SetWriteDeadline(time.Now().Add(time.Second * time.Duration(timeout)))
		_, err := conn.Write([]byte(req))
		if err != nil {
			fmt.Println(err)
			continue
		}
		buf := make([]byte, 0, 4096) // big buffer
		tmp := make([]byte, 256)     // using small tmo buffer for demonstrating
		for {
			n, err := conn.Read(tmp)
			if err != nil {
				if err != io.EOF {
					fmt.Println("read error:", err)
				}
				break
			}
			//fmt.Println("got", n, "bytes.")
			buf = append(buf, tmp[:n]...)
		}
		fmt.Println(string(buf))
		// var result string
		// for {
		// 	var data []byte = make([]byte, 1024)
		// 	n, err := conn.Read(data)
		// 	if n == 0 || err != nil {
		// 		break
		// 	}
		// 	result += string(data)
		// }
		// fmt.Println(result)
		time.Sleep(1 * time.Second)
	}
}

func StartServer() {
	message := "Hello, I am a server" // отправляемое сообщение
	listener, err := net.Listen("tcp", ":4545")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.Write([]byte(message))
		conn.Close()
	}
}
