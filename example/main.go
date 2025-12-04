package main

import (
	"fmt"
	"github.com/steel-feel/gozec"
)

func main() {
	wallet, err := gozec.Init("../../new_wallet", gozec.Testnet)
	if err != nil {
		panic("Wallet init failed")
	}

	addresses := wallet.GetAddress()

	fmt.Printf("wallet addresses are\n Transaparent - %v\n Unified - %v\n", addresses.TAddress, addresses.UAddress ) 

}