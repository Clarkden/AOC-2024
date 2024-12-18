package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(inputFile)

	column1 := []int{}
	column2 := []int{}

	occurances := make(map[int]int)

	for scanner.Scan() {

		values := strings.Split(scanner.Text(), "   ")

		val1 := 0
		val2 := 0

		val1, err = strconv.Atoi(values[0])
		if err != nil {
			log.Fatal(err)
		}

		val2, err = strconv.Atoi(values[1])
		if err != nil {
			log.Fatal(err)
		}

		column1 = append(column1, val1)
		column2 = append(column2, val2)

		occurances[val1] = 0
	}

	for _, v := range column2 {
		if _, found := occurances[v]; found {
			if found {
				occurances[v]++
			}
		}
	}

	sort.Ints(column1)
	sort.Ints(column2)

	distance := 0
	similarity := 0

	for i, v := range column1 {
		if v < column2[i] {
			distance += column2[i] - v
			continue
		}

		distance += v - column2[i]
	}

	for k, v := range occurances {
		similarity += k * v
	}

	log.Printf("Distance: %d", distance)
	log.Printf("Similarity: %d", similarity)
}
