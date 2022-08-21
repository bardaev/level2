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
	var timeout int = *flag.Int("timeout", 10, "Timeout in seconds")
	flag.Parse()

	var host string = flag.Arg(0)
	var port string = flag.Arg(1)
	// var host string = "217.65.3.21"
	// var port string = "80"

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), time.Duration(timeout)*time.Second)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer conn.Close()

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		var req string = sc.Text()
		resRq := req + " / HTTP/1.0\r\n\r\n"
		_, err := conn.Write([]byte(resRq))
		if err != nil {
			fmt.Println(err)
			continue
		}

		go func() {
			buf := bufio.NewReader(conn)
			for {
				resp, err := buf.ReadString('\n')
				if err == io.EOF {
					break
				}
				if err != nil {
					continue
				}
				fmt.Print(resp)
			}
		}()
	}
}
