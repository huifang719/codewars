package deadfishswim

func DeadfishSwim(data string) []int {
	var result []int
	var value int
	for _, v := range data {
		switch v {
		case 'i':
			value++
		case 'd':
			value--
		case 's':
			value *= value
		case 'o':
			result = append(result, value)
		}
	}
	return result
}