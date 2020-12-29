package gameutils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func Input(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	str, _ := reader.ReadString('\n')
	return str[:len(str)-1]
}

func BoolInput(prompt string) bool {
	input := strings.ToLower(Input(prompt))
	if strings.HasPrefix(input, "y") {
		return true
	} else if strings.HasPrefix(input, "n") {
		return false
	} else {
		fmt.Println("Invalid input.")
		return BoolInput(prompt)
	}
}

func Shuffle(slice []string) {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

func Pop(slice []string) (string, []string) {
	item, newSlice := slice[0], slice[1:]
	return item, newSlice
}