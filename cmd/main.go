package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type InputLog struct {
//	date time.Time
	command string
	value string
}

func NewInputLog(command string, value string) *InputLog {
	i := new(InputLog)
	i.command = command
	i.value = value
	return i
}

var InputLogArray []InputLog

func main() {
	file, err := os.Open(`test.txt`)
	if err != nil {
		// Open error
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for i := 1; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			// error
			break;
		}
		if strings.Contains(sc.Text(), "InputLog") {
			s := strings.Split(sc.Text(), " ")[6]
			command := strings.Split(s, ":")[0]
			value := strings.Split(s, ":")[1]
			InputLogArray = append(InputLogArray, *NewInputLog(command, value))
		}
	}
	fmt.Println(InputLogArray)
}
