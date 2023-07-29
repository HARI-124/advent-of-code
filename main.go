package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Intersection(a []string) (c string) {

	set1 := make(map[string]bool)
	set2 := make(map[string]bool)
	set3 := make(map[string]bool)

	for _, item := range a[0] {

		set1[string(item)] = true
	}
	for _, item := range a[1] {

		set2[string(item)] = true
	}
	for _, item := range a[2] {

		set3[string(item)] = true
	}

	for char := range set1 {

		if set2[char] && set3[char] {
			c = char
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

	content, err := ioutil.ReadAll(file)

	lines := strings.Split(string(content), "\n")

	paragraphs := make([]string, 0)
	for i := 0; i < len(lines); i += 3 {
		end := i + 3
		if end > len(lines) {
			end = len(lines)
		}
		paragraph := strings.Join(lines[i:end], "\n")
		paragraphs = append(paragraphs, paragraph)
	}
	score := 0

	for _, paragraph := range paragraphs {

		rack := strings.Split(paragraph, "\n")
		if len(rack) > 2 {

			common := Intersection(rack)
			fmt.Println(common)
			ascii := []rune(common)
			if int(ascii[0]) > 96 {
				// fmt.Println(int(ascii[0]) - 96)
				score = score + int(ascii[0]) - 96
			} else {

				// fmt.Println(int(ascii[0]) - 38)
				score = score + int(ascii[0]) - 38
			}
		}

	}

	fmt.Println(score)

}
