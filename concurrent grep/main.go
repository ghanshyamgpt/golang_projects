// main.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

func grepFile(filename string, pattern string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening %s: %s\n", filename, err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	lineNumber := 1

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Printf("Error reading %s: %s\n", filename, err)
			return
		}

		if strings.Contains(line, pattern) {
			fmt.Printf("%s:%d: %s", filename, lineNumber, line)
		}

		if err == io.EOF {
			break
		}

		lineNumber++
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: concurrentgrep PATTERN FILE1 [FILE2 FILE3 ...]")
		return
	}

	pattern := os.Args[1]
	files := os.Args[2:]

	var wg sync.WaitGroup

	for _, filename := range files {
		wg.Add(1)
		go grepFile(filename, pattern, &wg)
	}

	wg.Wait()
}
