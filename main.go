package main

import (
	"fmt"
	restAPI "github.com/godoio/godoMain/restAPI"
)

func main() {
	fmt.Println("This is the Go-do App!")
	restAPI.GetUI()
}