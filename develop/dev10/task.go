package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// go StartServer()
	var timeout int = *flag.Int("timeout", 10, "Timeout in seconds")
	flag.Parse()

	// var host string = flag.Arg(0)
	// var port string = flag.Arg(1)
	var host string = "217.65.3.21"
	var port string = "80"

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
		// var req string = "GET /"
		conn.SetWriteDeadline(time.Now().Add(time.Second * time.Duration(timeout)))
		_, err := conn.Write([]byte(req))
		if err != nil {
			fmt.Println(err)
			continue
		}

		buf, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(buf)
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
