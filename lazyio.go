package lazyio

import (
	"os"
	"bufio"
	"fmt"
	"log"
)

func Input(s ...string) string {
	if len(s) > 0 {
		for _, str := range s {
			fmt.Printf(str)
		}
	}

	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	return value
}

func Output(s ...string) {
	for _, str := range s {
		os.Stdout.WriteString(str)
	}
}

func ReadFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 64*1024), 64*1024) // Set buffer size for optimal performance

	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines, scanner.Err()

	/*
		Usage:
		lines, err := lazyio.ReadFile("file.txt")
		if err != nil {
			log.Fatal("Error reading file:", err)
		}

		for _, line := range lines {
			lazyio.Output(line)
		}
	*/
}
