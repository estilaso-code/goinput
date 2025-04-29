package goinput

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Input reads a line of text from stdin after showing a prompt
func Input(prompt string) (string, error) {
	fmt.Print(prompt)
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(text), nil
}

// InputInt reads a line of text and tries to parse it as an int
func InputInt(prompt string) (int, error) {
	text, err := Input(prompt)
	if err != nil {
		return 0, err
	}
	num, err := strconv.Atoi(text)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %w", err)
	}
	return num, nil
}

// InputBool reads a line of text and tries to parse it as a bool

func InputBool(prompt string) (bool, error) {
	text, err := Input(prompt)
	if err != nil {
		return false, err
	}
	b, err := strconv.ParseBool(text)
	if err != nil {
		return false, fmt.Errorf("invalid boolean: %w", err)
	}
	return b, nil
}
