package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	tarkNum := []int{0, 0}
	f, err := os.Open("day2_p1/data.txt")
	if err != nil {
		fmt.Println("could not open file for this reason: ", err)
		return
	}
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println("could not close file for this reason: ", err)
			return
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		counts := map[rune]int{}
		for _, x := range s.Text() {
			counts[x]++
		}
		var twos, threes bool
		for _, c := range counts {
			if c == 2 {
				twos = true
			}
			if c == 3 {
				threes = true
			}
		}
		if twos {
			tarkNum[0]++
		}
		if threes {
			tarkNum[1]++
		}
	}
	fmt.Println("tarcknum: ", tarkNum[0]*tarkNum[1])
}
