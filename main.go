package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Intersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

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
		rack := strings.Split(line, "")

		rack1 := rack[:len(rack)/2]
		rack2 := rack[len(rack)/2:]
		// fmt.Println("rack1 :", len(rack1), " rack2 :", len(rack2))
		common := Intersection(rack1, rack2)
		ascii := []rune(common[0])
		if int(ascii[0]) > 96 {
			fmt.Println(int(ascii[0]) - 96)
			score = score + int(ascii[0]) - 96
		} else {

			fmt.Println(int(ascii[0]) - 38)
			score = score + int(ascii[0]) - 38
		}

	}

	fmt.Println(score)

}
