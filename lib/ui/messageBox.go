package ui

import (
	"fmt"
	"strings"
)

func PrintMessage(message string, width int) error {
	if width <= 0 {
		return fmt.Errorf("ширина текста не должна быть меньше нуля")
	}
	words := strings.Fields(message)
	if len(words) == 0 {
		return fmt.Errorf("текст не должет быть пустым")
	}
	var currentLine string
	for _, word := range words {
		if len(currentLine)+len(word)+1 > width {
			fmt.Println(currentLine)
			currentLine = word
		} else {
			if len(currentLine) > 0 {
				currentLine += " "
			}
			currentLine += word
		}
	}
	if len(currentLine) > 0 {
		fmt.Println(currentLine)
	}

	return nil
}
func PrintQuote(message string, width int, lineColor string, textColor string) error {

	if width <= 0 {
		return fmt.Errorf("ширина текста не должна быть меньше нуля")
	}

	lines := strings.Split(message, "\n")

	for _, line := range lines {
		words := strings.Fields(line)
		if len(words) == 0 {
			fmt.Println(lineColor + "│ " + "\033[0m")
			continue
		}

		var currentLine string
		for _, word := range words {
			if len(currentLine)+len(word)+1 > width {
				fmt.Println(lineColor + "│ " + "\033[0m" + textColor + currentLine + "\033[0m")
				currentLine = word
			} else {
				if len(currentLine) > 0 {
					currentLine += " "
				}
				currentLine += word
			}
		}
		if len(currentLine) > 0 {
			fmt.Println(lineColor + "│ " + "\033[0m" + textColor + currentLine + "\033[0m")
		}
	}
	return nil
}
