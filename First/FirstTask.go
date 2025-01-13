package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		min_num := 9
		min_idx := 0
		input, _ = in.ReadString('\n')
		input = input[:len(input)-1]
		for j, el := range input {
			g := int(el) - 48
			if g <= min_num {
				min_num = g
				min_idx = j
			} else {
				break
			}
		}
		if len(input) != 1 {
			fmt.Println(input[:min_idx] + input[min_idx+1:])
		} else {
			fmt.Println(0)
		}
	}
}
