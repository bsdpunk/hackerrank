package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Meal struct {
	Price float64
	Tip   float64
	Tax   float64
	Final int64
}

func Round(f float64) float64 {
	return math.Floor(f + .5)
}
func main() {
	//var hasMoreInLine bool
	scanner := bufio.NewScanner(os.Stdin)
	var counter uint = 0
	var ds Meal
	//	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		words := string(scanner.Text())
		if counter == 0 {
			integ, err := strconv.ParseFloat(words, 64)
			ds.Price = integ
			if err != nil {
				//				return err
			}
		}
		if counter == 1 {
			integ, err := strconv.ParseFloat(words, 64)
			if err != nil {
				//				return err
			}
			ds.Tip = ds.Price * (integ / 100)
		}
		if counter == 2 {
			integ, err := strconv.ParseFloat(words, 64)
			if err != nil {
				//				return err
			}
			ds.Tax = ds.Price * (integ / 100)
		}
		counter++

	}
	ds.Final = int64(Round(ds.Tip + ds.Price + ds.Tax))
	//fmt.Printf("%+v\n", ds)
	fmt.Printf("The total meal cost is %d dollars.\n", ds.Final)
}
