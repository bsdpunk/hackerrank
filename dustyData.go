package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type DataStr struct {
	iTwo uint64
	dTwo float64
	sTwo string
}

func main() {
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "
	var counter uint = 0
	var ds DataStr
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		words := string(scanner.Text())
		//fmt.Println(reflect.TypeOf(words))
		if counter == 0 {
			integ, err := strconv.ParseUint(words, 10, 64)
			ds.iTwo = integ
			if err != nil {
				//				return err
			}
		}
		if counter == 1 {
			integ, err := strconv.ParseFloat(words, 64)
			if err != nil {
				//				return err
			}
			ds.dTwo = integ
		}
		if counter == 2 {
			//integ, err := strconv.ParseUint(words, 10, 64)
			ds.sTwo = words
		}
		counter++
		//pset, err := parseWords(words)
		//what, err := solveForX(pset)
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}
		//fmt.Println(what)
		fmt.Printf("%+v\n", ds)

		//fmt.Println(scanner.Text())
	}

	// Declare second integer, double, and String variables.

	// Read and save an integer, double, and String to your variables.

	// Print the sum of both integer variables on a new line.
	fmt.Println(i + ds.iTwo)

	// Print the sum of the double variables on a new line.
	//fmt.Println(d + ds.dTwo)
	fmt.Printf("%1.1f\n", d+ds.dTwo)
	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println(s + ds.sTwo)
}
