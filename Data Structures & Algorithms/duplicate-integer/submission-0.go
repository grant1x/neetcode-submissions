import "slices"

func hasDuplicate(nums []int) bool {

	var seen []int = []int{}

	for i := range nums {
		if slices.Contains(seen, nums[i]) {
			return true
		} else {
			seen = append(seen, nums[i])
		}
	}
	return false
}