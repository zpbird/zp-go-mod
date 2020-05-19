// Package zinput import ...
package zinput

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// GetNum ...
func GetNum() {

}

func Input(name string, regexpText string) (number int) {
	var validNumber = false
	for !validNumber {
		fmt.Println("Please input a", name, ": ")
		reader := bufio.NewReader(os.Stdin)
		inputBytes, _, err := reader.ReadLine()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Occur error when input", name, ":", err)
			continue
		}
		inputText := string(inputBytes)
		validNumber, err = regexp.MatchString(regexpText, inputText)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Occur error when match", name, "(", inputText, "):", err)
			continue
		}
		if validNumber {
			number, err = strconv.Atoi(inputText)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Occur error when convert", name, "(", inputText, "):", err)
				continue
			}
		} else {
			fmt.Fprintln(os.Stdout, "The", name, "(", inputText, ") does not have the correct format!")
		}
	}
	fmt.Println("The input", name, ": ", number)
	return
}
