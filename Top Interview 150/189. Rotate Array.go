func rotate(nums []int, k int)  {
	k = k % len(nums)
	aux := make([]int, k)
	copy(aux[:k], nums[len(nums)-k:])
	copy(nums[k:], nums[:len(nums)-k])
	copy(nums[:k], aux)
}
