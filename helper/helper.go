package helper

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadFile(name string) []string {
	data, err := os.Open(name)

	if err != nil {
		fmt.Println(err)
		return []string{}
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)
	content := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		content = append(content, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return content
}

func ReadFileByCharacters(name string) [][]string {
	data, err := os.Open(name)

	if err != nil {
		fmt.Println(err)
		return [][]string{{}}
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)
	content := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		charactersLine := strings.Split(line, "")
		content = append(content, charactersLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return content
}
