func merge(nums1 []int, m int, nums2 []int, n int) {
	var auxNums[]int
    for i:= range nums1{
        auxNums= append(auxNums, nums1[i])
    }
    var pointerOne, pointerTwo,b,a int 
    for i:=0;i<n+m;i++{
        if(pointerOne<m ){
            a= auxNums[pointerOne]
        }
        if(pointerTwo<n){
            b= nums2[pointerTwo]
        }
        
        if(pointerOne<m && a<b) || pointerTwo==n{
            nums1[i]=a
            pointerOne++
        }else{
            nums1[i]=b
            pointerTwo++
        }
    }
}