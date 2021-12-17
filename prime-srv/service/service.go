package service

func GetPrime(input int64) (ret []int64) {
	var i int64
	for i = 0; i < input; i++ {
		if isPrime(i) {
			ret = append(ret, i)
		}
	}
	return
}

func isPrime(v int64) bool {
	var i int64 = 2
	for i = 2; i < v/2; i++ {
		if v%i == 0 {
			return false
		}
	}
	return true
}
