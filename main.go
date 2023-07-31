package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func FindUnique(a, b, c, d byte) bool {
	occurences := make(map[byte]bool)
	occurences[a] = true
	occurences[b] = true
	occurences[c] = true
	occurences[d] = true

	if len(occurences) == 4 {
		return true
	}

	return false

}

func Iterate(a string) {

	for i := 0; i < len(a); i++ {

		// end := i+14
		result := FindUnique(a[i], a[i+1], a[i+2], a[i+3])
		if result {
			fmt.Println(i + 3)
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
