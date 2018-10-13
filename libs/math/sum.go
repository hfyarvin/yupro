package math

// 从startNum加到endNum
func SumOfStartToEnd(start, end int64) (sum int64) {
	sum = 0
	for index := start; index < end+1; index++ {
		sum += index
	}
	
	return sum
}