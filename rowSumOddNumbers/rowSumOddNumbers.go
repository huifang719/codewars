package rowsumoddnumbers

func RowSumOddNumbers(n int) int {
	// get the first number of the row
	// first number 1+(n-1)*n 1=> 1, 2=> 3, 3=> 7, 4=> 12, 5=> 21
	firstNumber := 1 + (n-1)*n	

	sum := firstNumber // initialize sum with first number
	// get the next number of the row
	// next number is the previous number + 2
	for i := 1; i < n; i++ {
		firstNumber += 2
		sum += firstNumber
	}

	return sum
}

