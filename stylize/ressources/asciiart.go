package ascii

import (
	"os"
	"strings"
)

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

	if len(input) == 0 {
		return ""
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

func printLine(inputLine string, slice [][]string) string {
	var result string

	// checks if the char is not printable
	for _, char := range inputLine {
		if char < 32 || char > 126 {
			return "Error: Characters must be ASCII printable.\nFor reference on what constitutes printable characters, please consult this image: 'https://upload.wikimedia.org/wikipedia/commons/1/1b/ASCII-Table-wide.svg'.\nPrintable characters are those with decimal values ranging from 32 to 126."

		}
	}

	for j := 0; j < 8; j++ {
		for _, char := range inputLine {
			index := int(char) - 32
			result += slice[index][j]
		}
		result += "\n"
	}

	return result
}
