package str

func GetChar(s string, pos int) string {
	return string([]rune(s)[pos])
}
