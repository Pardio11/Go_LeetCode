func removeDuplicates(nums []int) int {
    var existing [205]bool
    var move int
    for i:=range nums{
        if(existing[nums[i]+100]==false){
            existing[nums[i]+100]=true
            nums[i-move]=nums[i];
        }else{
            move++;
        }
    }
    return len(nums)-move
}