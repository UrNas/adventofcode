package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("could not open file for this reason: ", err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println("could not close file for this reason: ", err)
		}
	}()
	s := bufio.NewScanner(f)
	fmt.Println("sum: ", sum(s))
}
func sum(s *bufio.Scanner) int {
	var n = 0
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			continue
		}
		n += i
	}
	return n
}
