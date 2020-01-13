package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main () {
	fmt.Println("Welcome to go tic tac toe.  There's no turning back now...")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Take your next move (X, Y)")

		text, _ := reader.ReadString('\n')
		fmt.Println("Got it,", text)

		if strings.TrimSpace(text) == "quit" {
			break
		}
	}
}
