package main

import "fmt"

type purchases struct {
	purchase []int64
	percentOfCashback int64
	cashback int64

}

func (arg purchases) countCashback() int64  {
	var totallll int64

	for _, value := range arg.purchase {
		totallll += value
	}
	var cashback int64
	cashback = arg.percentOfCashback * totallll / 100
	return cashback

}

func main() {

	a := purchases{
		purchase:          []int64{122,152,1,20,12},
		percentOfCashback: 15,
		cashback:          0,
	}

	a.cashback = a.countCashback()
	fmt.Println(a.cashback)
}
