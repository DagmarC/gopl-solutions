package main

import (
	"bufio"
	"fmt"
	"strings"
)

type wordCounter int
type lineCounter int

func (w *wordCounter) Write(p []byte) (int, error) {
	*w = wordCounter(count(string(p), bufio.ScanWords))
	return len(p), nil
}

func (l *lineCounter) Write(p []byte) (int, error) {
	*l = lineCounter(count(string(p), bufio.ScanLines))
	return len(p), nil
}

func main() {

	var wordsCount wordCounter
	wordsCount.Write([]byte("Only three words ad3\ndHAHA\nNewline"))
	fmt.Println("Count of words:", wordsCount)

	var linesCounter lineCounter
	linesCounter.Write([]byte("Only three words ad3\ndHAHA\nNewline"))
	fmt.Println("Count of lines:", linesCounter)
}

func count(input string, fn bufio.SplitFunc) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(fn)
	count := 0 // according to split fn
	for scanner.Scan() {
		count++
	}
	return count
}
