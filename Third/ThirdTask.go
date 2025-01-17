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

func removeDupSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

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
		if numElement == 0 {
			continue
		}

		input, _ = in.ReadString('\n')
		input = strings.TrimSpace(input)
		input3 := strings.Split(input, " ")
		intInput := make([]int, 0)
		for _, el := range input3 {
			buff, _ := strconv.Atoi(el)
			intInput = append(intInput, buff)
		}
		sort.Ints(intInput)

		//Валидация строки
		input1, _ := in.ReadString('\n')
		input1 = input1[:len(input1)-1]
		clean := removeDupSpaces(input1)
		clean = strings.TrimSpace(clean)
		if clean == "" {
			fmt.Println("No")
			continue
		}
		if clean[0] == 48 && len(clean) > 1 {
			fmt.Println("No")
			continue
		}
		if len(clean) != len(input1) || input1[0] == ' ' || input1[len(input1)-1] == ' ' {
			fmt.Println("No")
			continue
		}
		if clean[0] == '-' {
			if clean[1] == 48 && len(clean) > 2 {
				fmt.Println("No")
				continue
			}
		}
		input1 = strings.TrimSpace(input1)
		input2 := strings.Split(input1, " ")
		intInput2 := make([]int, 0)
		for _, el := range input2 {
			if el[0] == '-' {
				if el[1] == 48 && len(el) > 2 {
					fmt.Println("No")
					goto next
				}
			}
			if el[0] == 48 && len(el) > 1 {
				fmt.Println("No")
				goto next
			}
			buff, err := strconv.Atoi(el)
			if err != nil {
				break
			}
			intInput2 = append(intInput2, buff)
		}
		if numElement != len(input2) {
			fmt.Println("No")
			continue
		}
		if reflect.DeepEqual(intInput, intInput2) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	next:
	}
}
