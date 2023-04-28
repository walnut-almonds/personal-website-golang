package constant

import "regexp"

const (
	PasswordMaxLen = 25
	PasswordMinLen = 6
)

var CheckPasswordRegex = regexp.MustCompile(`[^a-zA-Z\d]`)
