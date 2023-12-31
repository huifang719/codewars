package reverseWord

func ReverseWord(s string) string {
	var result string
	var word string
	for _, v := range s {
		if v == ' ' {
			result += word + string(v)
			word = ""
		} else {
			word = string(v) + word
		}
	}
	result += word
	return result
}
