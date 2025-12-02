package gozec

import (
	"errors"
	"os"
	"unsafe"
)

/*
#cgo LDFLAGS: -L${SRCDIR}/ffi -lrust_ffi_go
#include <stdlib.h>
void go_create_wallet(const char* str, uint32_t network );
void go_sync(const char* str);
void go_get_txn_list(const char* ptr, const char* uuid);
typedef struct { char* uuid; char* uivk; char* ufvk; char* source; } CAccount;
typedef struct { char* t_address; char* u_address; } CAddress;
typedef struct { CAccount* ptr; size_t len; } CAccountArray;
typedef struct { char* height ; uint64_t total ; uint64_t orchard ;uint64_t unshielded ; } CBalance;
CAccountArray go_list_accounts(const char* str);
CAddress go_get_address(const char* ptr, const char* uuid);
char* go_send_txn(const char* wallet_name, const char* uuid, const char* address,uint64_t value, size_t target_note_count, uint64_t min_split_output_value, const char* memo  );
CBalance go_balance(const char* ptr, const char* uuid);

void free_struct_array(CAccountArray);
void free_string(const char* s);
*/
import "C"

var (
	instance *GozecWallet
)

type ZecAccount struct {
	uuid string
	uivk string
	ufvk string
}

type ZecAddress struct {
	tAddress string
	uAddress string
}

type GozecWallet struct {
	walletDir string
	account   ZecAccount
}

type ZecBalance struct {
	unshielded uint64
	shielded   uint64
	total      uint64
	height     string
}

type NetworkType int

// Define constants for the days of the week
const (
	Testnet NetworkType = iota
	Mainnet
)

func Init(wallet_dir string, networkType NetworkType) (*GozecWallet, error) {
	f, err := os.Open(wallet_dir)
	c_wallet_dir := C.CString(wallet_dir)

	defer C.free(unsafe.Pointer(c_wallet_dir))
	if err != nil {
		if os.IsNotExist(err) {
			//create the wallet
			c_network := C.uint32_t(networkType)
			C.go_create_wallet(c_wallet_dir, c_network)
		} else {
			return nil, errors.New("Unknown error")
		}
	}
	defer f.Close()

	accountCList := C.go_list_accounts(c_wallet_dir)
	defer C.free_struct_array(accountCList)

	accountList := (*[1 << 28]C.CAccount)(unsafe.Pointer(accountCList.ptr))[:accountCList.len:accountCList.len]

	instance = &GozecWallet{
		walletDir: wallet_dir,
		account: ZecAccount{
			uuid: C.GoString(accountList[0].uuid),
			uivk: C.GoString(accountList[0].uivk),
			ufvk: C.GoString(accountList[0].ufvk),
		},
	}

	return instance, nil
}

/*
Get Transparent and Shielded address of zcash account
*/
func (g *GozecWallet) GetAddress() ZecAddress {
	c_wallet_dir := C.CString(g.walletDir)
	c_uuid := C.CString(g.account.uuid)

	defer C.free(unsafe.Pointer(c_wallet_dir))
	defer C.free(unsafe.Pointer(c_uuid))

	C_accAddress := C.go_get_address(c_wallet_dir, c_uuid)

	// C_accAddress.t_address
	return ZecAddress{
		tAddress: C.GoString(C_accAddress.t_address),
		uAddress: C.GoString(C_accAddress.u_address),
	}

}

/*
Sync's balance of zcash account
*/
func (g *GozecWallet) Sync() {
	c_wallet_dir := C.CString(g.walletDir)

	defer C.free(unsafe.Pointer(c_wallet_dir))

	C.go_sync(c_wallet_dir)
}
/*
Get balances of wallet
*/
func (g *GozecWallet) GetBalance() ZecBalance {
	c_wallet_dir := C.CString(g.walletDir)
	c_uuid := C.CString(g.account.uuid)

	defer C.free(unsafe.Pointer(c_wallet_dir))
	defer C.free(unsafe.Pointer(c_uuid))

	C.go_sync(c_wallet_dir)

	balances := C.go_balance(c_wallet_dir, c_uuid)

	return ZecBalance{
		height:     C.GoString(balances.height),
		shielded:   uint64(balances.orchard),
		unshielded: uint64(balances.unshielded),
		total:      uint64(balances.total),
	}
}

/*
@param to : Could be transparent address (t....) or unified(u.....) address

@param value : Amount to send
*/
func (g *GozecWallet) SendTransaction(to string, value uint64) string {
	c_wallet_dir := C.CString(g.walletDir)
	c_uuid := C.CString(g.account.uuid)
	c_to := C.CString(to)

	defer C.free(unsafe.Pointer(c_wallet_dir))
	defer C.free(unsafe.Pointer(c_uuid))
	defer C.free(unsafe.Pointer(c_to))

	c_value := C.uint64_t(value)

	C_txn := C.go_send_txn(c_wallet_dir, c_uuid, c_to, c_value, C.uintptr_t(0), C.uint64_t(0), C.CString(""))
	defer C.free_string(C_txn)

	return C.GoString(C_txn)
}
