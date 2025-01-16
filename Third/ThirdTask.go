package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	input, _ := in.ReadString('\n')
	input = input[:len(input)-1]
	num, _ := strconv.Atoi(input)

	for i := 0; i < num; i++ {
		input, _ = in.ReadString('\n')
		input = strings.TrimSpace(input)
		numElement, _ := strconv.Atoi(input)

		sortinput := make([]string, len(input))
		for j := 0; j < numElement; j++ {
			input, _ = in.ReadString('\n')
			input = strings.TrimSpace(input)
			input := strings.Split(input, " ") // input - int
			fmt.Println(input)
			copy(sortinput, input)
			sort.Strings(sortinput)
		}
		for j := 0; j < numElement; j++ {
			input, _ = in.ReadString('\n')
			input = strings.TrimSpace(input)
			input := strings.Split(input, " ") // input - int
			fmt.Println(input)
			if reflect.DeepEqual(input, sortinput) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}
