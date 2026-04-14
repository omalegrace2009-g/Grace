package main

import (
	"strconv"
	"strings"
	"regexp"
)

// CONVERTS HEXADECIMAL TO A DECIMAL
func HexToDec(text string) string {
	words := strings.Fields(text)
	var result []string
	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			hexVal := words[i-1]
			hex, err := strconv.ParseInt(hexVal, 16, 64)
			if err == nil {
				result[len(result)-1] = strconv.FormatInt(hex, 10)
			}
			continue
		}
		result = append(result, words[i])
	}
	return strings.Join(result, " ")
}

// CONVERTS BINARY TO A DECIMAL
func BinToDec(text string) string {
	words := strings.Fields(text)
	var result []string
	for i := 0; i < len(words); i++ {
		if words[i] == "(bin)" && i > 0 {
			binVal := words[i-1]
			bin, err := strconv.ParseInt(binVal, 2, 64)
			if err == nil {
				result[len(result)-1] = strconv.FormatInt(bin, 10)
			}
			continue
		}
		result = append(result, words[i])
	}
	return strings.Join(result, " ")
}

// HANDLES (UP) & (UP, N), (LOW) & (LOW, N), (CAP) & (CAP, N)
func CaseTransForm(text string) string {
	words := strings.Fields(text)
	var result []string
	for i := 0; i < len(words); i++ {
		d := words[i]
		if strings.HasPrefix(d, "(") {
			for !strings.HasSuffix(d, ")") {
				i++
				d += " " + words[i]
			}
			p := strings.Split(strings.Trim(d, "()"), ",")
			act := strings.TrimSpace(p[0])
			n, _ := strconv.Atoi(strings.TrimSpace(strings.Join(p[1:], "")))
			if n < 1 {
				n = 1
			}
			for g := len(result) - n; g < len(result); g++ {
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

// HANDLES ARTICLE
func FixArticle(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words)-1; i++ {
		d := strings.ToLower(words[i+1])
		f := strings.ContainsRune("AEIOUHaeiouh", rune(d[0]))
		if strings.ToLower(words[i]) == "an" && f {
			if words[i] == "An" {
				words[i] = "A"
			} else {
				words[i] = "a"
			}
		}
		if strings.ToLower(words[i]) == "a" && f {
			if words[i] == "A" {
				words[i] = "An"
			} else {
				words[i] = "an"
			}
		}
	}
	return strings.Join(words, " ")
}

// HANDLES QUOTES
func FixQuote(text string) string {
	re := regexp.MustCompile(`'\s*([^']+?)\s*'`)    
	return re.ReplaceAllString(text, " '$1'")
}