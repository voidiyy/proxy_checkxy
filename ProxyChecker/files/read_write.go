package files

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Error opening file: %s", err)
		return nil, err
	}
	defer file.Close()

	var proxies []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			proxies = append(proxies, line)
		}

	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error scanning file: %s", err)
		return nil, err
	}

	return proxies, nil
}

func WriteToFile(path string, proxies []string) error {
	file, err := os.Create(path)
	if err != nil {
		log.Printf("Error opening file: %s", err)
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, proxy := range proxies {
		_, err := writer.WriteString(proxy + "\n")
		if err != nil {
			log.Printf("Error writing to file: %s", err)
			return err
		}
	}
	writer.Flush()
	return nil
}
