package main

import (
	"bufio"
	"fmt"
	"os"
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
		input, _ := in.ReadString('\n')
		input = input[:len(input)-1]
		input3 := strings.Split(input, " ")
		inputInt := make([]int, 2)
		inputInt[0], _ = strconv.Atoi(input3[0])
		inputInt[1], _ = strconv.Atoi(input3[1])
		sliceTable := make([][]rune, inputInt[0])
		posA := make([]int, 0)
		posB := make([]int, 0)
		for j := 0; j < inputInt[0]; j++ {
			input, _ = in.ReadString('\n')
			indexA := strings.IndexRune(input, 'A')
			indexB := strings.IndexRune(input, 'B')
			if indexA != -1 {
				posA = append(posA, j, indexA, j+indexA)
			}
			if indexB != -1 {
				posB = append(posB, j, indexB, j+indexB)
			}
			sliceTable[j] = []rune(input)
		}
		if posA[2] <= posB[2] {
			for posA[2] > 0 {
				if posA[1] > 0 {
					if sliceTable[posA[0]][posA[1]-1] == '.' {
						sliceTable[posA[0]][posA[1]-1] = 'a'
						posA[1]--
						posA[2]--
						continue
					}
				}
				if posA[0] > 0 {
					if sliceTable[posA[0]-1][posA[1]] == '.' {
						sliceTable[posA[0]-1][posA[1]] = 'a'
						posA[0]--
						posA[2]--
						continue
					}
				}
			}
			for posB[2] < inputInt[0]+inputInt[1]-2 {
				if posB[1] < inputInt[1]-1 {
					if sliceTable[posB[0]][posB[1]+1] == '.' {
						sliceTable[posB[0]][posB[1]+1] = 'b'
						posB[1]++
						posB[2]++
						continue
					}
				}
				if posB[0] < inputInt[0]-1 {
					if sliceTable[posB[0]+1][posB[1]] == '.' {
						sliceTable[posB[0]+1][posB[1]] = 'b'
						posB[0]++
						posB[2]++
						continue
					}
				}
			}
		} else {
			for posB[2] > 0 {
				if posB[1] > 0 {
					if sliceTable[posB[0]][posB[1]-1] == '.' {
						sliceTable[posB[0]][posB[1]-1] = 'b'
						posB[1]--
						posB[2]--
						continue
					}
				}
				if posB[0] > 0 {
					if sliceTable[posB[0]-1][posB[1]] == '.' {
						sliceTable[posB[0]-1][posB[1]] = 'b'
						posB[0]--
						posB[2]--
						continue
					}
				}
			}
			for posA[2] < inputInt[0]+inputInt[1]-2 {
				if posA[1] < inputInt[1]-1 {
					if sliceTable[posA[0]][posA[1]+1] == '.' {
						sliceTable[posA[0]][posA[1]+1] = 'a'
						posA[1]++
						posA[2]++
						continue
					}
				}
				if posA[0] < inputInt[0]-1 {
					if sliceTable[posA[0]+1][posA[1]] == '.' {
						sliceTable[posA[0]+1][posA[1]] = 'a'
						posA[0]++
						posA[2]++
						continue
					}
				}
			}
		}
		for _, row := range sliceTable {
			fmt.Fprint(out, string(row))
		}
	}
}
