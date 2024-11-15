package main

import "fmt"

func twosum(nums []int, target int) []int {

	nmap := make(map[int]int)

	for i, num := range nums {

		c := target - num

		if j, found := nmap[c]; found {
			return []int{i, j}
		}
		nmap[num] = i
	}
	return nil
}

func main() {
	array := []int{11, 16, 9, 13, 21, 3}
	sum := 24

	fmt.Println(twosum(array, sum))

}
