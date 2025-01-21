package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Directory struct {
	Dir     string   `json:"dir"`
	Files   []string `json:"files"`
	Folders []Folder `json:"folders"`
}

type Folder struct {
	Dir     string   `json:"dir"`
	Files   []string `json:"files"`
	Folders []Folder `json:"folders"`
}

func countHacked(dir Directory, flag bool) int {
	count := 0
	safe := flag
	for _, el := range dir.Files {
		if strings.HasSuffix(el, ".hack") {
			count += len(dir.Files)
			safe = false
			break
		}

	}
	for _, el := range dir.Folders {
		count += countHackedFolders(el, safe)
	}
	return count
}

func countHackedFolders(folder Folder, flag bool) int {
	count := 0
	safe := flag
	if flag {
		for _, el := range folder.Files {
			if strings.HasSuffix(el, ".hack") {
				count += len(folder.Files)
				safe = false
				break
			}
		}
		for _, el := range folder.Folders {
			count += countHackedFolders(el, safe)
		}
	} else {
		count += len(folder.Files)
		for _, el := range folder.Folders {
			count += countHackedFolders(el, safe)
		}
	}
	return count
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
		var jsonData strings.Builder
		for j := 0; j < numElement; j++ {
			line, _ := in.ReadString('\n')
			jsonData.WriteString(line)
		}

		decoder := json.NewDecoder(strings.NewReader(jsonData.String()))
		var dir Directory
		err := decoder.Decode(&dir)
		if err != nil {
			continue
		}
		fmt.Println(countHacked(dir, true))
	}

}
