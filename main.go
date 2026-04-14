package main 

import (
	"fmt"
	"os"
)
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	data, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	text := string(data)
	text = HexToDec(text)
	text = BinToDec(text)
	text = CaseTransForm(text)
	text = FixArticle(text)
	text = FixQuote(text)
	text = fixPunctuation(text)
	err = os.WriteFile(outputFile, []byte(text + "\n"), 0644)
	if err != nil {
		panic(err)
	}
}