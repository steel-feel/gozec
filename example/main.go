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

	balances := wallet.GetBalance()

	fmt.Printf("Balances are\n Total - %v\n UnShielded - %v\nShielded - %v\n Sapling %v\n",
	 balances.Total,
	 balances.Unshielded, 
	 balances.Shielded,
	 balances.Sapling,
	) 

	bob := "utest1zu25404davj828zv0d3uwsdtvtuxyqq4xzn07zuwxcgp74qtkym4ugrgn63ptf9h9z3wk8sqcqfp3xfs88ssaaufusj52p2rl8u7p3ukjk35k4e7thk72kgpf3pfp2t92pcdwjtgffnugjdpaheqhvmexgy0wdsv469h29937tfen9rss0nhpn9qyyxtmsmrt3c0thvlg6mhgyp6hc8"
	amount := uint64(10)
	fmt.Printf("Sending \n %v ZEC\n to %v\n", amount, bob)

	txn := wallet.SendTransaction(bob, amount )

	fmt.Printf("Transaction hash : %v \n", txn)


}