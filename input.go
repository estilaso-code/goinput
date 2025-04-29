package goinput

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Input reads a line of text from stdin after showing a prompt
func Input(prompt string) string {
	fmt.Print(prompt)
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimSpace(text)
}

// InputInt reads a line of text and tries to parse it as an int
func InputInt(prompt string) int {
	for {
		text := Input(prompt)
		if num, err := strconv.Atoi(text); err == nil {
			return num
		}
		fmt.Println("Enter a valid number, bro.")
	}
}

// InputBool reads a line of text and tries to parse it as a bool
func InputBool(prompt string) bool {
	for {
		text := Input(prompt)
		if b, err := strconv.ParseBool(text); err == nil {
			return b
		}
		fmt.Println("Type 'true' or 'false', bro.")
	}
}
