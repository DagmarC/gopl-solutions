package main

import "fmt"

func main() {
	fmt.Println(doReturn())
}

func doReturn() (res string) {
	defer func() {
		if r := recover(); r != nil {
			res = "I panicked. " + r.(string)
		}
	}()
	res = "I dont panic."
	panic("Whaat I am dying...")
}
