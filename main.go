package main

import (
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}
	input := os.Args[1]
	output := os.Args[2]
	data, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	text := string(data)
	text = HexToDec(text)
	text = BinToDec(text)
	text = CaseTransform(text)
	text = FixArticle(text)
	text = FixQuote(text)
	text = FixPunctuation(text)
err = os.WriteFile(output, []byte(text + "\n"), 0644)
if err != nil {
	panic(err)
}
}
