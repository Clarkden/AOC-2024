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

	rules := make(map[string][]string)
	updates := [][]string{}

	parsingRules := true
	for scanner.Scan() {

		inputStr := scanner.Text()

		if inputStr == "" {
			parsingRules = false
		}

		if parsingRules {
			ruleParts := strings.Split(inputStr, "|")

			rules[ruleParts[0]] = append(rules[ruleParts[0]], ruleParts[1])
		} else {
			updateArr := strings.Split(inputStr, ",")

			updates = append(updates, updateArr)
		}
	}

	total := 0
	for _, update := range updates {
		if isValidUpdate(update, rules) {
			midVal, _ := strconv.Atoi(update[len(update)/2])
			total += midVal
		}
	}
	log.Println("Sum of middle pages:", total)
}

func isValidUpdate(update []string, rules map[string][]string) bool {
	index := make(map[string]int)
	for i, page := range update {
		index[page] = i
	}
	for x, ys := range rules {
		if ix, ok := index[x]; ok {
			for _, y := range ys {
				if iy, ok := index[y]; ok && ix >= iy {
					return false
				}
			}
		}
	}
	return true
}
