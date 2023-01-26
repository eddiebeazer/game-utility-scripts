package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("cmd", "/C", "p4 clients -E jenkins-* > clientlist")
	err := cmd.Run()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}

	words, err := scanWords("clientlist")
	if err != nil {
		panic(err)
	}

	for _, word := range words {
		if strings.Contains(word, "jenkins-") {
			cmd := exec.Command("cmd", "/C", "p4 client -d -f "+word)
			err := cmd.Run()
			if err != nil {
				log.Printf("Command finished with error: %v", err)
			}

		}
	}
}

func scanWords(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	var words []string

	for scanner.Scan() {
		words = append(words, strings.Replace(scanner.Text(), "\x00", "", -1))
	}

	return words, nil
}
