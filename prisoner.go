package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var hasMoreInLine bool
	bio := bufio.NewReader(os.Stdin)
	line, hasMoreInLine, err := bio.ReadLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if hasMoreInLine != false {
		fmt.Println(err)
		os.Exit(1)
	}
	words := strings.Fields(string(line))
	fmt.Println(words)

}
