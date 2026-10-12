package p4077

func maxPrimes(n int, s int) []int {
	var primes []int
	lpf := make([]int, n+1)
	var sum int
	for i := 2; i <= n; i++ {
		if lpf[i] == 0 {
			sum += i
			if sum > s {
				break
			}
			primes = append(primes, i)
			lpf[i] = i
		}
		for _, j := range primes {
			if i*j > n {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}
	return primes
}
