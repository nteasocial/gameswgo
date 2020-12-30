package gameutils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Input(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	str, _ := reader.ReadString('\n')
	return str[:len(str)-1]
}

func Bool(input string) bool {
	if strings.HasPrefix(input, "y") {
		return true
	} else if strings.HasPrefix(input, "n") {
		return false
	} else {
		panic(fmt.Errorf("could not convert %s to bool", input))
	}
}

func BoolInput(prompt string) (b bool) {
	input := strings.ToLower(Input(prompt))
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("Invalid input.")
			b = BoolInput(prompt)
		}
	}()
	return Bool(input)
}

func Print(format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Printf(format, args...)
	time.Sleep(time.Millisecond * 500)
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
