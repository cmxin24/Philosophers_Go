package main

import (
	"fmt"
	"os"
)

func main() {
	argv, err := check_argv(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
