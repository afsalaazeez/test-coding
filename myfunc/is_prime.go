package myfunc

// IsPrimeIter function
func IsPrimeIter(val int) bool {
	if val <= 1 {
		return false
	}
	for i := 2; i < val; i++ {
		if val%i == 0 {
			return false
		}
	}
	return true
}

// IsPrimeRecur function
func IsPrimeRecur(val int) bool {
	return isPrime(val, val/2)
}

func isPrime(val int, i int) bool {
	if val <= 1 {
		return false
	} else if i == 1 {
		return true
	} else if val%i == 0 {
		return false
	} else {
		return isPrime(val, i-1)
	}
}
