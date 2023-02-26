package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	MegaByteLimit int `yaml:"mb_limit"`
	ReplaceTasks  []struct {
		PathRegex   string `yaml:"path"`
		SearchRegex string `yaml:"search"`
		AddAfter    string `yaml:"after"`
		AddBefore   string `yaml:"before"`
		Remove      bool   `yaml:"remove"`
	} `yaml:"replaces"`
}

func main() {
	loadYaml()
}

func loadYaml() {
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("failed to read YAML file: %v", err)
	}

	// unmarshal the YAML data into a Config struct
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("failed to unmarshal YAML data: %v", err)
	}

	fmt.Printf("File size limit: %d\n", config.MegaByteLimit)
	for i := 0; i < len(config.ReplaceTasks); i++ {
		// print the configuration
		var task = config.ReplaceTasks[i]
		fmt.Printf("Path Regex: %s\n", task.PathRegex)
		fmt.Printf("After: %s\n", task.AddAfter)
		fmt.Printf("Remove: %b\n", task.Remove)

	}
}

func icontain(search string, src string) bool {
	return strings.Contains(strings.ToLower(search), strings.ToLower(src))
}

func openFile() {
	// open the file for reading and writing
	file, err := os.OpenFile("file.txt", os.O_RDONLY, 0644)

	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}

	defer file.Close()

	// create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// read the file line by line
	for scanner.Scan() {
		line := scanner.Text()

		// check if the line contains the target text
		if icontain(line, "123") {
			// add new text after the target text
			newLine := line + " new text"
			_, err = fmt.Fprintln(file, newLine)
			if err != nil {
				fmt.Println("error writing new line:", err)
				return
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error reading file:", err)
	}
}
