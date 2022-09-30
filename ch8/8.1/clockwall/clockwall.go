package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

// RUN: ./clockwall NewYork=localhost:8010 Tokyo=localhost:8020
// while this is runnning: TZ=US/Eastern ./clockclient -port 8010 & TZ=Asia/Tokyo ./clockclient -port 8020
func main() {
	var wg sync.WaitGroup

	for _, arg := range os.Args[1:] {
		tz, addressP, err := parseTzPort(arg)
		if err != nil {
			log.Fatal(err)
		}
		wg.Add(1)
		go getTime(tz, addressP, wg)
	}
	wg.Wait()
}

func getTime(tz string, addressP string, wg sync.WaitGroup) {
	conn, err := net.Dial("tcp", addressP)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("TIMEZONE: ", tz)
	mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
	wg.Done()
}

func parseTzPort(arg string) (string, string, error) {
	res := strings.Split(arg, "=")
	if len(res) != 2 {
		return "", "", errors.New("invalid arg, expected 1 equals sign '=' in format tz=address:port")
	}
	return res[0], res[1], nil
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
