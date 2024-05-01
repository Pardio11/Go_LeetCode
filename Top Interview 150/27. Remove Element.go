func removeElement(nums []int, val int) int {
    var move int
    for i:= range nums{
        
        if(nums[i] == val){
            move++
            nums[i]=0
        }else{
            var aux int = nums[i]
            nums[i-move]=aux
        }
    }
    
    return len(nums)-move;
}