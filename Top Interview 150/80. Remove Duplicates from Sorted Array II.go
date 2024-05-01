func removeDuplicates(nums []int) int {
    var existing [205]int
    var move int
    for i:=range nums{
        if(existing[nums[i]+100]!=2){
            existing[nums[i]+100]++
            nums[i-move]=nums[i];
        }else{
            move++;
        }
    }
    return len(nums)-move
}