package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	buffer := bufio.NewScanner(file)

	reports := [][]int{}

	i := 0
	for buffer.Scan() {
		row := buffer.Text()

		strVals := strings.Split(row, " ")

		vals := []int{}

		for _, v := range strVals {
			val, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}

			vals = append(vals, val)
		}
		reports = append(reports, vals)

		i++
	}

	safeReports := 0

	for _, v := range reports {

		safe := true
		decreasing := true

		for x, val := range v {
			if len(v) == x+1 {
				continue
			}

			if x == 0 {
				if val < v[x+1] {
					decreasing = false
				}
			}

			if val < v[x+1] && decreasing {
				safe = false
				continue
			} else if val > v[x+1] && !decreasing {
				safe = false
				continue
			}

			delta := math.Abs(float64(val - v[x+1]))
			if delta < 1 || delta > 3 {
				safe = false
				continue
			}
		}

		if safe {
			safeReports++
		}

	}
	log.Printf("Safe Reports: %d", safeReports)

}
