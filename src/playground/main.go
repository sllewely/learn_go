package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main () {
	fmt.Println("Welcome to go tic tac toe.  There's no turning back now...")

	reader := bufio.NewReader(os.Stdin)
	pattern := regexp.MustCompile(`(?P<x>[012]) (?P<y>[012])`)
	for {
		fmt.Println("Take your next move (X, Y)")

		text, _ := reader.ReadString('\n')
		pattern.FindStringSubmatch(text)
		fmt.Println("Got it,", text)

		if strings.TrimSpace(text) == "quit" {
			break
		}
	}
}
