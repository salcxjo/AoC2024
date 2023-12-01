package strings

import "strings"

var numbers = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func FindFirstDigit(text string, words ...bool) int {
	findWords := len(words) > 0 && words[0]
	length := len(text)
	for i := 0; i < length; i++ {
		c := rune(text[i])
		if c >= '0' && c <= '9' {
			return int(c - '0')
		}
		if !findWords {
			continue
		}
		tmp := text[i:min(i+5, length)]
		for k, v := range numbers {
			if strings.HasPrefix(tmp, k) {
				return v
			}
		}
	}
	panic("no digit found")
}

func FindLastDigit(text string, words ...bool) int {
	findWords := len(words) > 0 && words[0]
	length := len(text)
	for i := length - 1; i >= 0; i-- {
		c := rune(text[i])
		if c >= '0' && c <= '9' {
			return int(c - '0')
		}
		if !findWords {
			continue
		}
		tmp := text[max(0, i-5) : i+1]
		for k, v := range numbers {
			if strings.HasSuffix(tmp, k) {
				return v
			}
		}
	}
	panic("no digit found")
}
