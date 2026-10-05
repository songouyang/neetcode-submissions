func twoSum(nums []int, target int) []int {
    m := make(map[int]int, len(nums))
	for index, num := range nums {
		if i, exist := m[target-num]; exist {
			return []int{i, index}
		}
		m[num] = index
	}
	return nil
}
