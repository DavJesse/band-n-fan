package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Too Many Arguments")
		return
	}
}
