package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
		if strings.Index(sc.Text(), "XXXXX") != -1 {
			fmt.Printf("%4d行目: %s\n", i, sc.Text())			
		}
	}
}
