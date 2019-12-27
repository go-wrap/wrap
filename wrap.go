package wrap

import "strings"

// Wrap wraps the given string forcefully with the given length.
func Wrap(s string, n int) string {
	if i := strings.IndexByte(s, '\n'); i >= 0 {
		return Wrap(s[:i], n) + "\n" + Wrap(s[i+1:], n)
	}
	if len(s) <= n {
		return s
	}
	return s[:n] + "\n" + Wrap(s[n:], n)
}

// At wraps the given string at the given byte. If possible, the string will be
// wrapped at the first appearance of the given byte before the given length.
// If the given byte only appears after the given length, the string will be
// wrapped at the first appearance. The string will be unaltered otherwise.
func At(s string, c byte, n int) string {
	if i := strings.IndexByte(s, '\n'); i >= 0 {
		return At(s[:i], c, n) + "\n" + At(s[i+1:], c, n)
	}
	if len(s) <= n {
		return s
	}
	if i := strings.LastIndexByte(s[:n], c); 0 <= i && i <= n {
		return s[:i] + "\n" + At(s[i+1:], c, n)
	}
	if i := strings.IndexByte(s, c); i >= 0 {
		return s[:i] + "\n" + At(s[i+1:], c, n)
	}
	return s
}

// Space wraps the given string at a whitespace character.
func Space(s string, n int) string { return At(s, ' ', n) }
