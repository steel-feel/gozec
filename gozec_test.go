package gozec

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
    zecWallet, err := Init("./hello")
    if err != nil {
        fmt.Printf("error %v", err.Error())
    }

    // fmt.Printf("uuid %v \nuivk %v\n ufvk %v\n", zecWallet.account.uuid, zecWallet.account.uivk, zecWallet.account.ufvk)

    if len(zecWallet.account.uuid) == 0 {
         t.Errorf("lend of uuid = 0 , want gt 0")
    }

}

func TestGetAddress(t *testing.T) {
    zecWallet, err := Init("./hello")
    if err != nil {
        fmt.Printf("error %v", err.Error())
    }

    accountAddress := zecWallet.getAddress()
   // fmt.Printf("Account address %v \n", accountAddress)
    if len(accountAddress.tAddress) == 0 {
         t.Errorf("tAddress should not be empty")
    }

}
