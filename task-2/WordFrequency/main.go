package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"unicode"
)

func splitIntoWords(input string) []string {
	input = strings.ToLower(input)
	cleaned := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return r
		}
		return ' '
	}, input)
	words := strings.Fields(cleaned)
	return words
}

func wordFrequency(input string) map[string]int {
	frequency := make(map[string]int)
	words := splitIntoWords(input)
	for _,word := range words{
		frequency[word]++
		
	}
	return frequency
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	input, _ := reader.ReadString('\n')

	frequency := wordFrequency(input)
	fmt.Println("Word Frequency:")
	for word, count := range frequency {
		fmt.Printf("%s: %d\n", word, count)
	}
}