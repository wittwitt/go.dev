package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	splitAmount := decimal.NewFromInt(80000)
	newPaymentAmount := decimal.NewFromInt(160808)

	itemAmount := newPaymentAmount

	splitIndex := 0
	for itemAmount.Cmp(splitAmount) > 0 {

		fmt.Println("i", splitIndex, itemAmount)

		itemAmount = itemAmount.Sub(splitAmount)
		splitIndex = splitIndex + 1
	}
}
