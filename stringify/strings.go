package stringify

func InterlacedString(first, second string) string {
	result, rest, shortest := "", "", 0

	if len(first) < len(second) {
		shortest = len(first)
		rest = string(second[:shortest])
	} else {
		shortest = len(second)
		rest = string(first[:shortest])
	}

	for i := 0; i < shortest; i++ {
		result += string(first[i]) + string(second[i])
	}

	return result + rest
}
