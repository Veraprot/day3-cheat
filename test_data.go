package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	goFile, goErr := os.ReadFile("nums-go.txt")
	if goErr != nil {
		panic(goErr)
	}
	goStr := strings.Split(string(goFile), " ")

	jsFile, err := os.ReadFile("partnums-js.txt")
	if err != nil {
		panic(err)
	}
	jsStr := strings.Split(string(jsFile), " ")

	for i, char := range jsStr {
		if goStr[i] != char {
			fmt.Println(i, char)
			fmt.Println(goStr[i], goStr[i-1])
			return
		}
	}
}
