package main

import "fmt"

type OurPurchasesInAMonth struct {
	purchases         []int64
	cashback          int64
	percentOfCashback int64
	sum               int64
}

func (argument OurPurchasesInAMonth) CountCashbackAnd() int64 {

	var total int64

	for _, value := range argument.purchases {
		total += value
	}

	var cashback int64

	cashback = argument.percentOfCashback * total / 100

	return cashback

}

func (argument OurPurchasesInAMonth) SumIt() int64 {

	var SumOfThisMonth int64

	for _, value := range argument.purchases {
		SumOfThisMonth += value
	}

	return SumOfThisMonth

}

func main() {

	NewValue := OurPurchasesInAMonth{
		purchases:         []int64{122, 222, 333, 444, 555, 666},
		cashback:          0,
		percentOfCashback: 15,
		sum:               0,
	}

	NewValue.cashback = NewValue.CountCashbackAnd()

	NewValue.sum = NewValue.SumIt()

	fmt.Println(NewValue.cashback, "15% of purchases")
	fmt.Println(NewValue.sum, "Sum of Purchases")

}
