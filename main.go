package main

import "fmt"

func Init() error {
	return nil
}

func main() {
	err := Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("hello world")
}
