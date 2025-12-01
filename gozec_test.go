package gozec

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
    want := "Hello, from gozec!"

    zecWallet, err := Init("./hello")
    if err != nil {
        fmt.Printf("error %v", err.Error())
    }

    if got := zecWallet.Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}

func TestAccount(t *testing.T) {
    zecWallet, err := Init("./hello")
    if err != nil {
        fmt.Printf("error %v", err.Error())
    }

    // fmt.Printf("uuid %v \nuivk %v\n ufvk %v\n", zecWallet.account.uuid, zecWallet.account.uivk, zecWallet.account.ufvk)

    if len(zecWallet.account.uuid) == 0 {
         t.Errorf("lend of uuid = 0 , want gt 0")
    }

}
