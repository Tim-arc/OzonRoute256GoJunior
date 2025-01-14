package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// M 77
// R 82
// C 67
// D 68
//Проверим что R выполняется 1 раз)

func Foo(input string) bool {
	map_command := make(map[int]bool)
	for _, el := range input {
		el1 := int(el)
		if el1 == 77 && !map_command[el1] {
			map_command[el1] = true
			map_command[82] = false
			map_command[67] = false
			map_command[68] = false
			continue
		} else if el1 == 77 {
			return false
		}
		if val, ex := map_command[77]; el1 == 82 && val && ex && !map_command[el1] {
			map_command[el1] = true
			map_command[77] = true
			map_command[68] = false
			map_command[67] = false
			continue
		} else if el1 == 82 && (map_command[82] || !ex || !val) {
			return false
		}
		if val, ex := map_command[77]; el1 == 67 && ex && val && !map_command[67] {
			map_command[el1] = true
			map_command[77] = false
			map_command[82] = false
			continue
		} else if el1 == 67 && (map_command[67] || !ex || !val) {
			return false
		}
		if val, ex := map_command[77]; el1 == 68 && ex && val && !map_command[68] {
			map_command[el1] = true
			map_command[77] = false
			map_command[82] = false
		} else if el1 == 68 && (map_command[68] || !ex || !val) {
			return false
		}
	}
	if map_command[68] {
		return true
	}
	return false
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
	flag := false
	for i := 0; i < num; i++ {
		input, _ := in.ReadString('\n')
		input = input[:len(input)-1]
		flag = Foo(input)
		if flag {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
