package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: licensr list/[license-name]")
		os.Exit(0)
	}
	fmt.Println(os.Args[1])
}
