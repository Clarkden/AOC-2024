package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		lineParts := strings.Split(line, ": ")

		if len(lineParts) != 2 {
			log.Fatal("Could not properly parse")
		}

		targetValue, err := strconv.Atoi(lineParts[0])
		if err != nil {
			log.Fatal("Could not parse target value")
		}

		nums := strings.Split(lineParts[1], " ")

		numValues := make([]int, len(nums))
		for i, v := range nums {
			numValue, err := strconv.Atoi(strings.TrimSpace(v))
			if err != nil {
				log.Fatal("Could not parse number value: ", v)
			}
			numValues[i] = numValue
		}

		if checkCombination(numValues, targetValue, 0, 0) {
			sum += targetValue
		}
	}

	log.Print(sum)
}

func checkCombination(nums []int, target, index, current int) bool {
	if index == len(nums) {
		return current == target
	}

	if checkCombination(nums, target, index+1, current+nums[index]) {
		return true
	}

	if index == 0 {
		return checkCombination(nums, target, index+1, nums[index])
	} else {
		if checkCombination(nums, target, index+1, current*nums[index]) {
			return true
		}
	}

	return false
}
