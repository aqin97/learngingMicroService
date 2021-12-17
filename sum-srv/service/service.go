package service

func GetSum(input int64) (ret int64) {
	var i int64
	for i = 0; i < input; i++ {
		ret += i
	}
	return
}
