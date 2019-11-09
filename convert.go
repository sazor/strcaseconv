package strcaseconv

import "strings"

const (
	snakeDelimiter = '_'
	kebabDelimiter = '-'
	capitalDiff    = 'a' - 'A'
)

// SnakeToUpperCamel converts original snake_case string to UpperCamelCase
func SnakeToUpperCamel(original string) string {
	return toCamel(original, snakeDelimiter, true)
}

// SnakeToLowerCamel converts original snake_case string to lowerCamelCase
func SnakeToLowerCamel(original string) string {
	return toCamel(original, snakeDelimiter, false)
}

// KebabToUpperCamel converts original kebab-case string to UpperCamelCase
func KebabToUpperCamel(original string) string {
	return toCamel(original, kebabDelimiter, true)
}

// KebabToLowerCamel converts original kebab-case string to lowerCamelCase
func KebabToLowerCamel(original string) string {
	return toCamel(original, kebabDelimiter, false)
}

// CamelToSnake converts original UpperCamelCase or lowerCamelCase string to snake_case
func CamelToSnake(original string) string {
	return fromCamel(original, snakeDelimiter)
}

// CamelToKebab converts original UpperCamelCase or lowerCamelCase string to kebab-case
func CamelToKebab(original string) string {
	return fromCamel(original, kebabDelimiter)
}

// SnakeToKebab converts original snake_case string to kebab-case
func SnakeToKebab(original string) string {
	return betweenSnakeKebab(original, snakeDelimiter, kebabDelimiter)
}

// KebabToSnake converts original kebab-case string to snake_case
func KebabToSnake(original string) string {
	return betweenSnakeKebab(original, kebabDelimiter, snakeDelimiter)
}

func fromCamel(original string, delimiter rune) string {
	if original == "" {
		return original
	}
	upperCamel := isUpper(original[0])
	var cnt int
	for i := 1; i < len(original); i++ {
		if isUpper(original[i]) {
			cnt++
		}
	}
	if cnt == 0 {
		if upperCamel {
			var kebabized strings.Builder
			kebabized.Grow(len(original))
			kebabized.WriteByte(original[0] + capitalDiff)
			kebabized.WriteString(original[1:])
			return kebabized.String()
		}
		return original
	}
	var kebabized strings.Builder
	kebabized.Grow(len(original) + cnt)
	if upperCamel {
		kebabized.WriteByte(original[0] + capitalDiff)
	} else {
		kebabized.WriteByte(original[0])
	}
	for i := 1; i < len(original); i++ {
		if isUpper(original[i]) {
			if isLower(original[i-1]) || isUpper(original[i-1]) && i+1 < len(original) && isLower(original[i+1]) {
				kebabized.WriteRune(delimiter)
			}
			kebabized.WriteByte(original[i] + capitalDiff)
		} else {
			kebabized.WriteByte(original[i])
		}
	}
	return kebabized.String()
}

func toCamel(original string, delimiter rune, upper bool) string {
	if original == "" {
		return original
	}
	var cnt int
	for i := 0; i < len(original); i++ {
		if original[i] == byte(delimiter) {
			cnt++
		}
	}
	if cnt == 0 && !upper {
		return original
	}
	var camelized strings.Builder
	camelized.Grow(len(original) - cnt)
	delimiterFound := upper
	for i := range original {
		if original[i] == byte(delimiter) {
			delimiterFound = true
			continue
		}
		c := original[i]
		if delimiterFound && isLower(c) {
			c -= capitalDiff
		}
		camelized.WriteByte(c)
		delimiterFound = false
	}
	return camelized.String()
}

func betweenSnakeKebab(original string, srcDelimiter, dstDelimiter rune) string {
	if original == "" || !strings.ContainsRune(original, srcDelimiter) {
		return original
	}
	var kebabized strings.Builder
	kebabized.Grow(len(original))
	for i := range original {
		if original[i] == byte(srcDelimiter) {
			kebabized.WriteRune(dstDelimiter)
		} else {
			kebabized.WriteByte(original[i])
		}
	}
	return kebabized.String()
}

func isLower(char byte) bool {
	return 'a' <= char && char <= 'z'
}

func isUpper(char byte) bool {
	return 'A' <= char && char <= 'Z'
}
