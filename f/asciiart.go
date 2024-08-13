package ascii

import (
	"fmt"
	"os"
	"strings"
)

func GenerateAsciiArt(input, banner string) (string, error) {
	var slice [][]string

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		return "", fmt.Errorf("Error: Not a valid banner")
	}

	content, err := os.ReadFile(banner + ".txt")
	if err != nil {
		return "", fmt.Errorf("Error reading file: %v", err)
	}

	str := string(content)
	var lines []string
	if banner != "thinkertoy" {
		lines = strings.Split(str, "\n\n")
	} else {
		fixedContent := strings.ReplaceAll(str, "\r\n", "\n")
		lines = strings.Split(fixedContent, "\n\n")
	}

	for i := range lines {
		if lines[i] != "" {
			if i == 0 {
				lines[0] = lines[0][1:]
			}
			slice = append(slice, strings.Split(lines[i], "\n"))
		}
	}

	if input == "" {
		return "", nil
	}

	var result strings.Builder

	inputLines := strings.Split(input, "\\n")
	for _, value := range inputLines {
		if value == "" {
			result.WriteString("\n")
		} else {
			asciiLine, err := printLine(value, slice)
			if err != nil {
				return "", err
			}
			result.WriteString(asciiLine)
		}
	}

	return result.String(), nil
}

func printLine(inputLine string, slice [][]string) (string, error) {
	var result strings.Builder

	// checks if the char is not printable
	for _, char := range inputLine {
		if char < 32 || char > 126 {
			return "", fmt.Errorf("Character '%c' is not a printable ASCII character", char)
		}
	}

	for j := 0; j < 8; j++ {
		for _, char := range inputLine {
			index := int(char) - 32
			result.WriteString(slice[index][j])
		}
		result.WriteString("\n")
	}

	return result.String(), nil
}
