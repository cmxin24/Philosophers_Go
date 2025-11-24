package main

import (
	"fmt"
	"os"
)

func main() {
	augs, err := checkArgv(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
