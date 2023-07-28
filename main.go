package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("err: ", err)

	}
	content, err := ioutil.ReadAll(file)
	paragraphs := bytes.Split(content, []byte("\n\n"))
	sum := 0
	calories := []int{}

	for _, paragraph := range paragraphs {
		fmt.Println(string(paragraph) + "\n")
		paragraph1 := string(paragraph)
		number := strings.Split(paragraph1, "\n")
		fmt.Println("num", number)
		isum := 0

		for _, num := range number {
			num2, _ := strconv.Atoi(num)
			isum = isum + num2
			calories = append(calories, isum)
		}

	}
	sort.Sort(sort.Reverse(sort.IntSlice(calories)))
	sum = calories[0] + calories[1] + calories[2]
	fmt.Println("largest calorie :", sum)
}
