package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	matrix := [][]string{}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	i := 0

	for scanner.Scan() {
		row := scanner.Text()
		rowArr := strings.Split(row, "")
		matrix = append(matrix, rowArr)
		i++
	}

	appearances := 0

	for i, l := range matrix {
		for i2 := range l {

			str := ""

			if len(l)-i2 <= 3 {
				break
			}

			for i3 := i2; i3 < i2+4; i3++ {
				str += string(l[i3])
			}

			if str == "XMAS" || str == "SAMX" {
				appearances += 1
			}

		}

		for i2 := range l {

			if len(matrix)-i <= 3 {
				break
			}

			str := ""

			for r := i; r < i+4; r++ {
				str += string(matrix[r][i2])
			}

			if str == "XMAS" || str == "SAMX" {
				appearances += 1
			}

		}

		for i2 := range l {

			if len(matrix)-i <= 3 || len(l)-i2 <= 3 {
				break
			}

			str := ""

			x := 0
			for r := i; r < i+4; r++ {
				str += string(matrix[r][i2+x])
				x++
			}

			if str == "XMAS" || str == "SAMX" {
				appearances += 1
			}

		}

		for i2 := range l {

			if len(matrix)-i <= 3 || i2 < 3 {
				continue
			}

			str := ""

			x := 0
			for r := i; r < i+4; r++ {
				str += string(matrix[r][i2-x])
				x++
			}

			if str == "XMAS" || str == "SAMX" {
				appearances += 1
			}

		}
	}

	log.Print(appearances)
}
