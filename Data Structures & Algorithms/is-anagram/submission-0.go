func isAnagram(s string, t string) bool {
	m := make(map[rune]int, len(s))
	for _, ch := range s {
		m[ch]++
	}
	for _, ch := range t {
		m[ch]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
