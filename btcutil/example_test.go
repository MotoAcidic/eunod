package eunoutil_test

import (
	"fmt"
	"math"

	"github.com/MotoAcidic/eunod/eunoutil"
)

func ExampleAmount() {

	a := eunoutil.Amount(0)
	fmt.Println("Zero Satoshi:", a)

	a = eunoutil.Amount(1e8)
	fmt.Println("100,000,000 Satoshis:", a)

	a = eunoutil.Amount(1e5)
	fmt.Println("100,000 Satoshis:", a)
	// Output:
	// Zero Satoshi: 0 EUNO
	// 100,000,000 Satoshis: 1 EUNO
	// 100,000 Satoshis: 0.001 EUNO
}

func ExampleNewAmount() {
	amountOne, err := eunoutil.NewAmount(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountOne) //Output 1

	amountFraction, err := eunoutil.NewAmount(0.01234567)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountFraction) //Output 2

	amountZero, err := eunoutil.NewAmount(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountZero) //Output 3

	amountNaN, err := eunoutil.NewAmount(math.NaN())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountNaN) //Output 4

	// Output: 1 EUNO
	// 0.01234567 EUNO
	// 0 EUNO
	// invalid bitcoin amount
}

func ExampleAmount_unitConversions() {
	amount := eunoutil.Amount(44433322211100)

	fmt.Println("Satoshi to kEUNO:", amount.Format(eunoutil.AmountKiloEUNO))
	fmt.Println("Satoshi to EUNO:", amount)
	fmt.Println("Satoshi to MilliEUNO:", amount.Format(eunoutil.AmountMilliEUNO))
	fmt.Println("Satoshi to MicroEUNO:", amount.Format(eunoutil.AmountMicroEUNO))
	fmt.Println("Satoshi to Satoshi:", amount.Format(eunoutil.AmountSatoshi))

	// Output:
	// Satoshi to kEUNO: 444.333222111 kEUNO
	// Satoshi to EUNO: 444333.222111 EUNO
	// Satoshi to MilliEUNO: 444333222.111 mEUNO
	// Satoshi to MicroEUNO: 444333222111 μEUNO
	// Satoshi to Satoshi: 44433322211100 Satoshi
}
