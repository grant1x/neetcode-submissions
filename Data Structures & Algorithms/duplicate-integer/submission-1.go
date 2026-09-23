func hasDuplicate(nums []int) bool {

	seen := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		if _, exists := seen[num]; exists  {
			return true
		} else {
			seen[num] = struct{}{}
		}
	}
	return false
}