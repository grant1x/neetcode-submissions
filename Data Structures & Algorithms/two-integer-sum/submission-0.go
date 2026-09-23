func twoSum(nums []int, target int) []int {

	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if complementIndex, ok := seen[complement]; ok {
			return []int{complementIndex, i}
		}
		seen[num] = i
	}
	return nil
}