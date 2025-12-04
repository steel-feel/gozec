package example

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

	fmt.Printf("wallet address is %v", addresses.UAddress ) 

}