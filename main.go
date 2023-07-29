package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("err: ", err)

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")
		fmt.Println("values: ", values)
		i := values[0]
		j := values[1]
		// chance of getting draw
		if i == "A" && j == "X" {
			score = score + 3

		}
		if i == "B" && j == "Y" {

			score = score + 5
		}
		if i == "C" && j == "Z" {

			score = score + 7

		}

		if i == "A" && j == "Y" {
			score = score + 4

		}
		if i == "B" && j == "Z" {

			score = score + 9
		}
		if i == "C" && j == "X" {

			score = score + 2

		}

		if i == "A" && j == "Z" {
			score = score + 8

		}
		if i == "B" && j == "X" {

			score = score + 1
		}
		if i == "C" && j == "Y" {

			score = score + 6

		}

	}

	fmt.Println(score)

}
