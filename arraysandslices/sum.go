package arraysandslices

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	// length := len(numbers)
	// sums := make([]int, length)
	// for i, nums := range numbers {
	// 	sums[i] = Sum(nums)
	// }
	// return sums
	var sums []int
	for _, nums := range numbers {
		sums = append(sums, Sum(nums))
	}
	return sums
}

func SumAllTails(numbers ...[]int) []int {
	var sums []int
	for _, nums := range numbers {
		if len(nums) <= 0 {
			sums = append(sums, 0)
		} else {
			tails := nums[1:]
			sums = append(sums, Sum(tails))
		}
	}
	return sums
}
