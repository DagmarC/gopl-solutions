package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

type server struct {
	address, port, tz string
}

// RUN: ./clockwall NewYork=localhost:8010 Tokyo=localhost:8020
// while this is runnning: TZ=US/Eastern ./clock -port 8010 & TZ=Asia/Tokyo ./clock -port 8020
func main() {

	var wg sync.WaitGroup

	servers := make([]*server, 0, len(os.Args))
	for _, arg := range os.Args[1:] {
		server, err := parseServer(arg)
		if err != nil {
			log.Fatal(err)
		}
		servers = append(servers, &server)
	}
	// Clockwall header
	fmt.Printf("Timezone\tAdress:Port\t\tTime\n")
	for _, s := range servers {
		wg.Add(1)
		go clientTCP(s, &wg)
	}
	wg.Wait()
}

func clientTCP(server *server, wg *sync.WaitGroup) {

	conn, err := net.Dial("tcp", server.address+":"+server.port)
	if err != nil {
		log.Fatal(err)
	}
	connbuf := bufio.NewReader(conn)
	for {
		str, err := connbuf.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Printf("%-6s\t\t%6s:%-4s\t\t%9s", server.tz, server.address, server.port, str)
	}
}

func parseServer(arg string) (server, error) {
	res := strings.Split(arg, "=")
	if len(res) != 2 {
		return server{}, errors.New("invalid arg, expected 1 equals sign '=' in format tz=address:port")
	}
	addr, prt, err := parseAddrPrt(res[1])
	if err != nil {
		return server{}, err
	}
	return server{tz: res[0], address: addr, port: prt}, nil
}

func parseAddrPrt(s string) (string, string, error) {
	res := strings.Split(s, ":")
	if len(res) != 2 {
		return "", "", errors.New("invalid arg, expected 1 equals sign ':' in format address:port")
	}
	return res[0], res[1], nil
}
