package main

import (
	"fmt"
	"github.com/zolinz/protobuf-test/simplepb"
)

func main() {
	fmt.Println("hello Zoli")
	doSimple()

}

func doSimple() {
	sm := simplepb.Simple{
		Name: "Zolika",
		Age:  46,
	}

	fmt.Print(sm)
}
