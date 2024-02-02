package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args

	if len(os.Args) != 3 {
		fmt.Println("指令长度不正确")
	}

	old_file := arg[1]
	new_file := arg[2]

}
