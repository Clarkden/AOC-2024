package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	buffer := bufio.NewScanner(file)

	contents := ""

	for buffer.Scan() {
		contents += buffer.Text()
	}

	re := regexp.MustCompile(`mul\(\s*(0|[1-9]\d*)\s*,\s*(0|[1-9]\d*)\s*\)`)

	matches := re.FindAllString(contents, -1)

	total := 0

	for _, match := range matches {

		numbersStr := strings.Trim(strings.Trim(match, "/mul("), ")")
		numbers := strings.Split(numbersStr, ",")

		num1, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatal(err)
		}

		num2, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatal(err)
		}

		total += num1 * num2
	}

	log.Print(total)
}
