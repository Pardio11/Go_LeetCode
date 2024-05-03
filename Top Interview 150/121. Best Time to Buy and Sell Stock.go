func maxProfit(prices []int) int {
    highestProfit:=0
    lowestCost:=prices[0]
    var currentProfit int
    for _,price:=range prices{
        currentProfit=price-lowestCost
        if highestProfit<currentProfit {
            highestProfit = currentProfit
        }
        if price<lowestCost {
            lowestCost = price 
        }
    }
    return highestProfit
    
}