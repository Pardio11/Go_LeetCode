func jump(nums []int) int {
    bestJump:=0
    bestReach:=nums[0]
    jumps:=0
	for i:=0;i<len(nums);{
		if bestReach>=len(nums)-1 || bestJump == len(nums)-1 {
			i=len(nums)
			if bestJump != len(nums)-1{
				jumps++
			}
			
		}else{
		for j:=i+1;j<=i+nums[i] && j<len(nums);j++{
            if bestReach<nums[j]+j {
                bestJump=j
                bestReach=nums[j]+j
            }
        }
		jumps++
		i=bestJump
		}
		
	}
    return jumps
}