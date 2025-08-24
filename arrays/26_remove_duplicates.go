package arrays

func RemoveDuplicates(nums []int) int {
	cur := 0
	for i := 1; i < len(nums); i++ {
		if nums[cur] != nums[i] {
			cur++
			nums[cur] = nums[i]
		}
	}
	return cur + 1
}
