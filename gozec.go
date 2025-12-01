package gozec

import (
	"errors"
	"os"
	"unsafe"
)

/*
#cgo LDFLAGS: -L${SRCDIR}/ffi -lrust_ffi_go
#include <stdlib.h>
void go_create_wallet(const char* str);
void go_sync(const char* str);
void go_get_txn_list(const char* ptr, const char* uuid);
typedef struct { char* uuid; char* uivk; char* ufvk; char* source; } CAccount;
typedef struct { CAccount* ptr; size_t len; } CAccountArray;
typedef struct { char* height ; uint64_t total ; uint64_t orchard ;uint64_t unshielded ; } CBalance;
CAccountArray go_list_accounts(const char* str);
char* go_get_address(const char* ptr, const char* uuid);
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

type GozecWallet struct {
	walletDir string
	account   ZecAccount
}

func Init(wallet_dir string) (*GozecWallet, error) {
	f, err := os.Open(wallet_dir)
	c_wallet_dir := C.CString(wallet_dir)
	defer C.free(unsafe.Pointer(c_wallet_dir))
	if err != nil {
		if os.IsNotExist(err) {
			//create the wallet
			C.go_create_wallet(c_wallet_dir)
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

// func (g *GozecWallet) fetchAccounts() error {
// 	c_wallet_dir := C.CString(g.walletDir)
// 	defer C.free(unsafe.Pointer(c_wallet_dir))
// 	accountCList := C.go_list_accounts(c_wallet_dir)
// 	defer C.free_struct_array(accountCList)

// 	accountList := (*[1 << 28]C.CAccount)(unsafe.Pointer(accountCList.ptr))[:accountCList.len:accountCList.len]

// 	for _, s := range accountList {
// 		g.accounts = append(g.accounts, ZecAccount{
// 			uuid: C.GoString(s.uuid),
// 			uivk: C.GoString(s.uivk),
// 			ufvk: C.GoString(s.ufvk),
// 		})
// 	}

// 	return nil

// }

// Hello returns a greeting.
func (w *GozecWallet) Hello() string {
	return "Hello, from gozec!"
}
