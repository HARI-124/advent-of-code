package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func FindUnique(a string) bool {
	occurences := make(map[byte]bool)

	// for _, char := range a {
	// 	occurences[char] = true
	// }
	for i := 0; i < len(a); i++ {
		occurences[a[i]] = true
	}

	if len(occurences) == 14 {
		return true
	}

	return false

}

func Iterate(a string) {

	for i := 0; i < len(a); i++ {

		end := i + 14
		result := FindUnique(a[i:end])
		fmt.Println(a[i:end])

		if result {
			fmt.Println(i + 14)
			return
		}
	}

}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("err: ", err)

	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)

	lines := strings.Split(string(content), "\n")
	// fmt.Println("sample", len(lines))

	Iterate(lines[0])
}
