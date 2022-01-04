package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter list: ")
	scanner := bufio.NewScanner(os.Stdin)
	var inlist = "("
	i := 0
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		}
		i++
		inlist = inlist + "\n'" + scanner.Text() + "',"
	}
	inlist = strings.TrimSuffix(inlist,",") + "\n)"
	fmt.Println("Here are your results")
	fmt.Println("Total values scanned:", strconv.Itoa(i))
	fmt.Println( inlist )
}