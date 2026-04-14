package main

import (
	"strconv"
	"strings"
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
		f := words[i]
		if strings.HasPrefix(f, "(") {
			for !strings.HasSuffix(f, ")") {
				i++
				f += " " + words[i]
			}
			p := strings.Split(strings.Trim(f, "()"), ",")
			act := strings.TrimSpace(p[0])
			n, _ := strconv.Atoi(strings.TrimSpace(strings.Join(p[1:], "")))
			if n < 1 {
				n = 1
			}
			for d := len(result) - n; d < len(result); d++ {
				if d >= 0 {
					switch act {
					case "up":
						result[d] = strings.ToUpper(result[d])
					case "low":
						result[d] = strings.ToLower(result[d])
					case "cap":
						result[d] = strings.Title(result[d])
					}
				}
			}
			continue
		}
		result = append(result, f)
	}
	return strings.Join(result, " ")
}

// HANDLES ARTICLE
func FixArticle(text string) string {
	words := strings.Fields(text)
	var result []string
	for i := 0; i < len(words); i++ {
		if words[i] == "a" {
			words[i] += "an " + words[i]
		}
		if words[i] == "A" {
			words[i] += "AN " + words[i]
		}
		if words[i] == "an" {
			words[i] += "a " + words[i]
		}
		if words[i] == "AN" {
			words[i] += "A " + words[i]
		}
		result = append(result, words[i])
	}
	return strings.Join(result, " ")
}
