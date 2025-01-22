package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func AbraCodabra(sliceVans [][]int, sliceOrder []int, intInput []int, resultChan chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, el := range intInput {
		find := false
		minKey := 0
		for k, v := range sliceVans {
			if el <= v[1] && el >= v[0] {
				minKey = k
				find = true
				break
			}
			if el > v[0] && el < v[1] {
				break
			}
		}
		if find {
			sliceVans[minKey][2] -= 1
			for k, v := range sliceOrder {
				if v == el {
					sliceOrder[k] = sliceVans[minKey][3] + 1
					break
				}
			}
			if sliceVans[minKey][2] == 0 {
				sliceVans = append(sliceVans[:minKey], sliceVans[minKey+1:]...)
			}
		} else {
			for k, v := range sliceOrder {
				if v == el {
					sliceOrder[k] = -1
					break
				}
			}
		}
	}
	resultChan <- sliceOrder
}

func main() {
	var in *bufio.Reader
	in = bufio.NewReader(os.Stdin)

	input, _ := in.ReadString('\n')
	input = input[:len(input)-1]
	num, _ := strconv.Atoi(input)
	for i := 0; i < num; i++ {
		input, _ = in.ReadString('\n')
		input = strings.TrimSpace(input)
		numElement, _ := strconv.Atoi(input)

		input, _ = in.ReadString('\n')
		input = strings.TrimSpace(input)
		input3 := strings.Split(input, " ")
		sliceOrder := make([]int, len(input3))
		intInput := make([]int, 0)
		for i, el := range input3 {
			buff, _ := strconv.Atoi(el)
			sliceOrder[i] = buff
			intInput = append(intInput, buff)
		}
		sort.Ints(intInput)

		input, _ = in.ReadString('\n')
		input = strings.TrimSpace(input)
		numElement, _ = strconv.Atoi(input)

		//Считываем строки с грузовиками
		sliceVans := make([][]int, numElement)
		for i := 0; i < numElement; i++ {
			input, _ = in.ReadString('\n')
			input = strings.TrimSpace(input)
			input3 := strings.Split(input, " ")
			sliceVan := make([]int, len(input3)+1)
			for i, el := range input3 {
				buff, _ := strconv.Atoi(el)
				sliceVan[i] = buff
			}
			sliceVans[i] = sliceVan
			sliceVans[i][3] = i
		}
		sort.Slice(sliceVans, func(i, j int) bool {
			if sliceVans[i][0] == sliceVans[j][0] {
				return sliceVans[i][3] < sliceVans[j][3]
			}
			return sliceVans[i][0] < sliceVans[j][0]
		})

		resultChan := make(chan []int)
		var wg sync.WaitGroup
		wg.Add(1)
		go AbraCodabra(sliceVans, sliceOrder, intInput, resultChan, &wg)

		go func() {
			wg.Wait()
			close(resultChan)
		}()

		result := <-resultChan
		for _, v := range result {
			fmt.Print(v, " ")
		}
	}
}

// Решение через горутины

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"sort"
// 	"strconv"
// 	"strings"
// 	"sync"
// )

// func AbraCodabra(sliceVans [][]int, sliceOrder []int, intInput []int, resultChan chan []int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	var mutex sync.Mutex
// 	for _, el := range intInput {
// 		find := false
// 		minKey := 0
// 		for k, v := range sliceVans {
// 			if el <= v[1] && el >= v[0] {
// 				minKey = k
// 				find = true
// 				break
// 			}
// 			if el > v[0] && el < v[1] {
// 				break
// 			}
// 		}
// 		if find {
// 			sliceVans[minKey][2] -= 1
// 			for k, v := range sliceOrder {
// 				if v == el {
// 					sliceOrder[k] = sliceVans[minKey][3] + 1
// 					break
// 				}
// 			}
// 			if sliceVans[minKey][2] == 0 {
// 				sliceVans = append(sliceVans[:minKey], sliceVans[minKey+1:]...)
// 			}
// 		} else {
// 			for k, v := range sliceOrder {
// 				if v == el {
// 					sliceOrder[k] = -1
// 					break
// 				}
// 			}
// 		}
// 	}
// 	mutex.Lock()
// 	resultChan <- sliceOrder
// 	mutex.Unlock()
// }

// func main() {
// 	resultChan := make(chan []int)
// 	var wg sync.WaitGroup
// 	var in *bufio.Reader
// 	in = bufio.NewReader(os.Stdin)

// 	input, _ := in.ReadString('\n')
// 	input = input[:len(input)-1]
// 	num, _ := strconv.Atoi(input)
// 	for i := 0; i < num; i++ {
// 		input, _ = in.ReadString('\n')
// 		input = strings.TrimSpace(input)
// 		numElement, _ := strconv.Atoi(input)

// 		input, _ = in.ReadString('\n')
// 		input = strings.TrimSpace(input)
// 		input3 := strings.Split(input, " ")
// 		sliceOrder := make([]int, len(input3))
// 		intInput := make([]int, 0)
// 		for i, el := range input3 {
// 			buff, _ := strconv.Atoi(el)
// 			sliceOrder[i] = buff
// 			intInput = append(intInput, buff)
// 		}
// 		sort.Ints(intInput)

// 		input, _ = in.ReadString('\n')
// 		input = strings.TrimSpace(input)
// 		numElement, _ = strconv.Atoi(input)

// 		//Считываем строки с грузовиками
// 		sliceVans := make([][]int, numElement)
// 		for i := 0; i < numElement; i++ {
// 			input, _ = in.ReadString('\n')
// 			input = strings.TrimSpace(input)
// 			input3 := strings.Split(input, " ")
// 			sliceVan := make([]int, len(input3)+1)
// 			for i, el := range input3 {
// 				buff, _ := strconv.Atoi(el)
// 				sliceVan[i] = buff
// 			}
// 			sliceVans[i] = sliceVan
// 			sliceVans[i][3] = i
// 		}
// 		sort.Slice(sliceVans, func(i, j int) bool {
// 			if sliceVans[i][0] == sliceVans[j][0] {
// 				return sliceVans[i][3] < sliceVans[j][3]
// 			}
// 			return sliceVans[i][0] < sliceVans[j][0]
// 		})

// 		wg.Add(1)
// 		go AbraCodabra(sliceVans, sliceOrder, intInput, resultChan, &wg)

// 	}
// 	go func() {
// 		wg.Wait()
// 		close(resultChan)
// 	}()

// 	for result := range resultChan {
// 		for _, v := range result {
// 			fmt.Print(v, " ")
// 		}
// 		fmt.Println()
// 	}
// }
