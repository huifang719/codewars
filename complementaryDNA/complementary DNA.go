package complementaryDNA

func ComplementaryDNA(dna string) string {
	var result string
	for _, v := range dna {
		switch v {
		case 'A':
			result += "T"
		case 'T':
			result += "A"
		case 'C':
			result += "G"
		case 'G':
			result += "C"
		}
	}
	return result
}
