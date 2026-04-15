package main

import (
	"strings"
	"strconv"
	"regexp"
)
//HEXADECIMAL TO DECIMAL
func HexToDec(text string) string {
	word := strings.Fields(text)
	var result []string
	for i := 0; i < len(word); i++ {
		if word[i] == "(hex)" && i > 0 {
			hexVal := word[i -1]
			hex, err := strconv.ParseInt(hexVal, 16, 64)
			if err == nil {
				result[len(result) -1] = strconv.FormatInt(hex, 10)
			}
			continue
		}
		result = append(result, word[i])
	}
	return strings.Join(result, " ")
}

//BINARY TO DECIMAL
func BinToDec(text string) string {
	word := strings.Fields(text)
	var result []string
	for i := 0; i < len(word); i++ {
		if word[i] == "(bin)" && i > 0 {
			hexVal := word[i -1]
			hex, err := strconv.ParseInt(hexVal, 2, 64)
			if err == nil {
				result[len(result) -1] = strconv.FormatInt(hex, 10)
			}
			continue
		}
		result = append(result, word[i])
	}
	return strings.Join(result, " ")
}

//HANDLES (UP) & (UP ,N), (LOW) & (LOW, N), (CAP) & (CAP, N)
func CaseTransform(text string) string {
	word := strings.Fields(text)
	var result []string
	for i := 0; i < len(word); i++ {
		d := word[i]
		if strings.HasPrefix(d, "(") {
			for !strings.HasSuffix(d, ")") {
				i++
				d += " " + word[i]
			}
			p := strings.Split(strings.Trim(d, "()"), ",")
			act := strings.TrimSpace(p[0])
			n, _ := strconv.Atoi(strings.TrimSpace(strings.Join(p[1:], "")))
			if n < 1 {
				n = 1
			}
			for g := len(result)-n; g < len(result); g++ {
				if g >= 0 {
					switch act {
					case "up":
						result[g] = strings.ToUpper(result[g])
					case "low":
						result[g] = strings.ToLower(result[g])
					case "cap":
						result[g] = strings.Title(result[g])
					}
				}
			}
			continue
		}
		result = append(result, d)
	}
	return strings.Join(result, " ")
}

//HANDLE ARTICLE
func FixArticle(text string) string {
word := strings.Fields(text)
for i := 0; i < len(word)-1; i++ {
	s := strings.ToLower(word[i+1])
	d := strings.ContainsRune("aeiouhAEIOUH", rune(s[0]))
	if strings.ToLower(word[i]) == "an" && d {
		if word[i] == "An" {
			word[i] = "A"
		} else {
			word[i] = "a"
		}
	}
		if strings.ToLower(word[i]) == "a" && d {
			if word[i] == "A" {
				word[i] = "An"
			} else {
				word[i] = "an"
			}
		}
	}
return strings.Join(word, " ")
}
//HANDLES QUOTES
func FixQuote(text string) string {
s := regexp.MustCompile(`'\s*([^']+?)\s*'`)
return s.ReplaceAllString(text, "'$1'")
}

//HANDLE PUNCTUATION
func FixPunctuation(text string) string {
	text = regexp.MustCompile(`.\s*\.s*\.`).ReplaceAllString(text,  "...")
	text = regexp.MustCompile(`\s+([,.:;?1]+)`).ReplaceAllString(text, "$1")
	return text
}