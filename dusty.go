package main

import (
	"bufio"
	"fmt"
	//	"math"
	"os"
	//	"reflect"
	"strconv"
	"strings"
)

//The random number, add the sweets number, that's the serial number
//unless it's above the amount of serial numbers, then it's the difference
//between the interim_serial number, and the amount of serial numbers
type PrisonerSet struct {
	rando  int64
	sweets int64
	total  int64
}

func parseWords(p []string) (PrisonerSet, error) {

	rando, err := strconv.ParseInt(p[2], 10, 64)
	sweets, err := strconv.ParseInt(p[1], 10, 64)
	total, err := strconv.ParseInt(p[0], 10, 64)

	if err != nil {
		fmt.Println(err)
	}

	return PrisonerSet{rando: rando, sweets: sweets, total: total}, nil
}

func solveForX(pset PrisonerSet) (int64, error) {
	psion := pset.rando + pset.sweets
	double := pset.total * 2
	//var final int64 := psion
	if pset.sweets > double {
		//pset.sweets
		if psion%pset.total == psion {
			//			fmt.Println("1")
			return psion - 1, nil
		}
		if (psion-1)%pset.total == 0 {
			//			fmt.Println("2")
			return pset.total, nil
		}
		//if (psion % pset.total) < 0 {
		//			fmt.Println("3")
		//	return psion, nil
		//}

		//	}
		//	fmt.Println("4")
		//fmt.Println(psion)
		return (psion - 1) % pset.total, nil
	} else if psion > pset.total {

		psion := -1*(pset.total-psion) - 1
		if psion-1 == pset.total {
			return 1, nil
		}
		if psion == 0 {
			return pset.total, nil
		}

		if psion > pset.total {
			return psion - pset.total, nil
		}
		return psion, nil
	} else {
		return psion - 1, nil
	}
}

func main() {
	//var hasMoreInLine bool
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		words := strings.Fields(string(scanner.Text()))
		//fmt.Println(reflect.TypeOf(words))
		if len(words) < 2 {
			fmt.Printf("")
		} else {
			pset, err := parseWords(words)
			what, err := solveForX(pset)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(what)
			//fmt.Printf("%+v\n", pset)

		}
		//fmt.Println(scanner.Text())
	}

	//bio := bufio.NewReader(os.Stdin)
	//line, hasMoreInLine, err := bio.ReadLine()

	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//if hasMoreInLine != false {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//words := strings.Fields(string(line))
	//fmt.Println(reflect.TypeOf(words))
	//pset, err := parseWords(words)
	//what, err := solveForX(pset)
	//fmt.Println(what)
	//fmt.Printf("%+v\n", pset)

}
