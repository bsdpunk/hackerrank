package main

import "fmt"
import "os"
import "bufio"

func main() {
	//var inputString string
	//input := bufio.NewReader(os.Stdin)
	//inputString, err := input.ReadString('\n')
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

	//fmt.Scanln(&inputString)
	fmt.Println("Hello, World.")
	fmt.Println(string(line))
}
