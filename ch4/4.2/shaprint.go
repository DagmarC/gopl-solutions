package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	shaPtr := flag.Int("sha", 256, "Generate sha256/384/512 digest.")
	flag.Parse()

	word := os.Args[2]

	switch *shaPtr {
	case 256:
		fmt.Println(sha256.Sum256([]byte(word)))

	case 384:
		fmt.Println(sha512.Sum384([]byte(word)))

	case 512:
		fmt.Println(sha512.Sum512([]byte(word)))

	default:
		fmt.Println("Invalid hash format.")
	}
}
