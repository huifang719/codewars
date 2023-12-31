package rowsumoddnumbers

func RowSumOddNumbers(n int) int {
	// get the first number of the row
	// first number 1+(n-1)*n 1=> 1, 2=> 3, 3=> 7, 4=> 12, 5=> 21
	firstNumber := 1 + (n-1)*n	
	sum := firstNumber
	for i := 1; i <= n; i++ {
		numberAfter := firstNumber + 2*i
		sum += numberAfter
	}
	return n*n*n
}