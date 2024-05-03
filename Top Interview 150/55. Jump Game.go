func canJump(nums []int) bool {
    reachable:=true
    distance:=1
    for i:=len(nums)-2;i>=0;i--{
        if(distance <= nums[i]){
            distance=1
            reachable=true
        }else{
            reachable=false
            distance++
        }
    }
    return reachable
}