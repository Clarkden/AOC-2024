package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	mappedArea := [][]string{}

	for scanner.Scan() {

		inputStr := scanner.Text()

		inputArr := strings.Split(inputStr, "")

		mappedArea = append(mappedArea, inputArr)
	}

	guardX := 0
	guardY := 0
	direction := ""

	for i := 0; i < len(mappedArea); i++ {
		for j := 0; j < len(mappedArea[i]); j++ {
			currentChar := mappedArea[i][j]
			if currentChar == "v" || currentChar == ">" || currentChar == "<" || currentChar == "^" {

				switch currentChar {
				case "v":
					direction = "down"
				case ">":
					direction = "right"
				case "<":
					direction = "left"
				case "^":
					direction = "up"
				default:
					log.Fatal("Invalid character")
				}

				guardX = j
				guardY = i
				break
			}
		}
	}

	distinctMoves := 1
	inMappedArea := true
	visitedPairs := [][]int{}

	for inMappedArea {

		// log.Printf("Guard X: %d Guard Y: %d Direction: %s", guardX, guardY, direction)
		// log.Print("---------------------")
		// for i := 0; i < len(mappedArea); i++ {
		// 	log.Print(mappedArea[i])
		// }
		// log.Print("---------------------")

		switch direction {
		case "up":
			wantedMove := guardY - 1

			if wantedMove < 0 {
				inMappedArea = false
				break
			}

			if mappedArea[wantedMove][guardX] == "#" {
				direction = "right"
				mappedArea[guardY][guardX] = ">"
			} else {
				mappedArea[guardY][guardX] = "."
				mappedArea[wantedMove][guardX] = "^"
				guardY--
			}
		case "right":
			wantedMove := guardX + 1

			if wantedMove > len(mappedArea[guardY])-1 {
				inMappedArea = false
				break
			}

			if mappedArea[guardY][wantedMove] == "#" {
				direction = "down"
				mappedArea[guardY][guardX] = "v"
			} else {
				mappedArea[guardY][guardX] = "."
				mappedArea[guardY][wantedMove] = ">"
				guardX++
			}
		case "left":
			wantedMove := guardX - 1

			if wantedMove < 0 {
				inMappedArea = false
				break
			}

			if mappedArea[guardY][wantedMove] == "#" {
				direction = "up"
				mappedArea[guardY][guardX] = "^"
			} else {
				mappedArea[guardY][guardX] = "."
				mappedArea[guardY][wantedMove] = "<"
				guardX--
			}
		case "down":
			wantedMove := guardY + 1

			if wantedMove > len(mappedArea)-1 {
				inMappedArea = false
				break
			}

			if mappedArea[wantedMove][guardX] == "#" {
				direction = "left"
				mappedArea[guardY][guardX] = "<"
			} else {
				mappedArea[guardY][guardX] = "."
				mappedArea[wantedMove][guardX] = "v"
				guardY++
			}
		default:
			log.Fatal("Invalid guard state")
		}

		distinct := true

		for _, pair := range visitedPairs {
			if pair[0] == guardX && pair[1] == guardY {
				distinct = false
				break
			}
		}

		if distinct {
			visitedPairs = append(visitedPairs, []int{guardX, guardY})
			distinctMoves++
		}

	}

	log.Print(distinctMoves)
}
