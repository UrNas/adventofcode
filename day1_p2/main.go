package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	nums := map[int]bool{0: true}
	n := 0
	f, err := os.Open("day1_p2/data.txt")
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
	data := creatSlice(s)
	for {
		for _, x := range data {
			n += x
			if nums[n] {
				fmt.Println("num: ", n)
				return
			}
			nums[n] = true
		}
	}

}
func creatSlice(s *bufio.Scanner) []int {
	var data []int
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			fmt.Println("could not convert for this reason: ", err)
			return nil
		}
		data = append(data, i)
	}
	return data
}
