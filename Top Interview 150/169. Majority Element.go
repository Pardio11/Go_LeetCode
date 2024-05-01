func majorityElement(nums []int) int {
    register := make(map[int]int)
    var max, maxValue int
    for i:=range nums{
        register[nums[i]]=register[nums[i]]+1
        if(maxValue<register[nums[i]]){
            maxValue=register[nums[i]]
            max=nums[i]
        }
    }
    return max
}