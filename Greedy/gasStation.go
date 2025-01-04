//go:build ignore

package main

import "fmt"

func getTotal(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func canCompleteCircuit(gas, cost []int) int {
	if getTotal(gas) < getTotal(cost) {
		return -1
	}
	total, res := 0, 0
	for i := 0; i < len(gas); i++ {
		total = gas[i] - cost[i]
		if total < 0 {
			total = 0
			res = i + 1
		}
	}
	return res
}

func main() {
	gas1 := []int{1, 2, 3, 4}
	cost1 := []int{2, 2, 4, 1}
	fmt.Println(canCompleteCircuit(gas1, cost1))
	gas2 := []int{1, 2, 3}
	cost2 := []int{2, 3, 2}
	fmt.Println(canCompleteCircuit(gas2, cost2))
}
