package gogenutils

import (
	"regexp"
	"strings"
	"unicode"
)

var (
	fieldNameRE     = regexp.MustCompile("[^A-Za-z0-9]+")
	jsonFieldNameRE = regexp.MustCompile("[^a-z0-9_]+|_$")
)

// FieldName returns a go struct field name from a text string
func FieldName(s string) string {
	s = PascalCase(s)
	s = fieldNameRE.ReplaceAllString(s, "")
	return s
}

// JSONFieldName returns the lowercase, underscore se
func JSONFieldName(s string) string {
	s = SnakeCase(s)
	s = jsonFieldNameRE.ReplaceAllString(s, "")
	s = strings.Replace(s, "__", "_", -1)

	// Strip excess
	return s
}

// CamelCase returns the camelCase version of a text string
func CamelCase(s string) string {
	return camelAndPascalCase(s, true)
}

// PascalCase returns the PascalCase version of a text string
func PascalCase(s string) string {
	return camelAndPascalCase(s, false)
}

func camelAndPascalCase(s string, isCamel bool) string {

	var (
		prev    = ' '
		isFirst = true
	)

	s = strings.Map(
		func(r rune) rune {
			defer func() {
				prev = r
			}()

			if isFirst {
				isFirst = false
				if isCamel {
					return unicode.ToLower(r)
				} else {
					return unicode.ToTitle(r)
				}
			}

			if isSeparator(prev) {
				return unicode.ToTitle(r)
			}
			return r
		},
		s)

	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "_", "", -1)

	return s
}

// SnakeCase returns the snake_case version of a text string
func SnakeCase(s string) string {
	return strings.Map(
		func(r rune) rune {
			if isSeparator(r) {
				return '_'
			}
			return unicode.ToLower(r)
		},
		s)
}

func isSeparator(r rune) bool {
	return !unicode.IsLetter(r) && !unicode.IsNumber(r)
}

// PascalCaseToSnakeCase converts from PascalCase to snake_case
func PascalCaseToSnakeCase(s string) string {
	var result string
	var words []string
	var lastPos int
	rs := []rune(s)

	for i := 0; i < len(rs); i++ {
		if i > 0 && isWordBreak(rs, i) {
			// Scan ahead looking for an initialism
			if initialism := startsWithInitialism(s[lastPos:]); initialism != "" {
				words = append(words, initialism)

				// Move the pointer fowards
				i += len(initialism) - 1
				lastPos = i
				continue
			}

			// Not an initialism, so break the previous word
			words = append(words, s[lastPos:i])
			lastPos = i
		}
	}

	// append the last word
	if s[lastPos:] != "" {
		words = append(words, s[lastPos:])
	}

	for k, word := range words {
		if k > 0 {
			result += "_"
		}

		result += strings.ToLower(word)
	}

	return result
}

func isWordBreak(rs []rune, i int) bool {
	if unicode.IsUpper(rs[i]) {
		return true
	}

	if (i == 0 || !unicode.IsDigit(rs[i-1])) && unicode.IsDigit(rs[i]) {
		return true
	}

	return false
}

// startsWithInitialism returns the initialism if the given string begins with it
func startsWithInitialism(s string) string {
	// the longest initialism is 5 char, the shortest 2
	for i := 5; i > 0; i-- {
		if len(s) >= i && commonInitialisms[s[:i]] {
			return s[:i]
		}
	}

	// Is this probably an initialism?
	rs := []rune(s)
	var end int
	for end = 0; end < 5; end++ {
		if len(rs) <= end {
			break
		}

		if !isWordBreak(rs, end) {
			break
		}
	}

	if end > 2 {
		return string(rs[:end-1])
	}

	return ""
}

// commonInitialisms, taken from
// https://github.com/golang/lint/blob/206c0f020eba0f7fbcfbc467a5eb808037df2ed6/lint.go#L731
var commonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
}
