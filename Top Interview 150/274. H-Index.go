func hIndex(citations []int) int {
    maxH:=0
    good:=true
    for i:=len(citations);i>0;i--{
		fail:=0
        for j:=0;j<len(citations);j++{
            good=true
            if citations[j]<i{
                fail++
            }
            if fail>len(citations)-i {
                good=false
               break 
            }
        }
        if good {
            maxH=i
            break
        }
    }
    return maxH
}