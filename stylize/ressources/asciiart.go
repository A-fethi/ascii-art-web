package ascii

import (
	"os"
	"strings"
)

// GenerateAsciiArt creates ASCII art from the given input and banner file.
func GenerateAsciiArt(input []string, banner string) string {
	var slice [][]string

	content, err := os.ReadFile(banner + ".txt")
	if err != nil {
		return ""
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

	var result string

	for _, value := range input {
		if value == "" {
			result += "\n"
		} else {
			asciiLine := printLine(value, slice)
			result += asciiLine
		}
	}
	return result
}

// printLine generates a line of ASCII art for a given input line using the provided slice.
func printLine(inputLine string, slice [][]string) string {
	var result string

	for j := 0; j < 8; j++ {
		for _, char := range inputLine {
			index := int(char) - 32
			result += slice[index][j]
		}
		result += "\n"
	}

	return result
}
