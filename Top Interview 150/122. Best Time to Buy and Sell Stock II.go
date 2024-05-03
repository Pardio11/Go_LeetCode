func maxProfit(prices []int) int {
    currentProfit := 0
    currentLowest := prices[0]
    currentHighest := prices[0]
    for _,price := range prices {
        if price < currentHighest {
            if (currentHighest-currentLowest) >0 {
                currentProfit += (currentHighest-currentLowest)
            }
            currentLowest = price     
        }  
        currentHighest=price
    }
    if (currentHighest-currentLowest) >0 {
        currentProfit += (currentHighest-currentLowest)
    }
    return currentProfit
}