// https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
// An interface is two things: it is a set of methods, but it is also a type.
package main

import (
	"fmt"
)

//kontrak interface
type IUser interface {
	GetUser() LAS_T_USER
	GetPDWK() PDWK
	// CreateUser() LAS_T_USER
}

//implementation
func (user LAS_T_USER) GetUser() LAS_T_USER {
	return user
}

func (user LAS_T_USER) GetPDWK() PDWK {
	return user.PDWK
}

type PN string
type UserStatus string
type PDWK struct {
	limit          int64
	limit_kupedes  int64
	limit_kur      int64
	limit_briguna  int64
	limit_ritel    int64
	limit_pangan   int64
	limit_menengah int64
	limit_cashcoll int64
	limit_kpr      int64
}
type LAS_T_USER struct {
	NO_PEGAWAI      PN
	NAMA_PEGAWAI    string
	USER_NAME       string
	PASSWORD        string
	FID_ROLE        int64
	KODE_CABANG     string
	fid_JHOOFF      string
	status          UserStatus
	sbdesc_uker     string
	unit_kerja_uker string
	PDWK            PDWK
}

func GetUserDetails(user IUser) {
	fmt.Println(user.GetUser())
	fmt.Println(user.GetPDWK())
}

func main() {
	var adk LAS_T_USER
	adk.NO_PEGAWAI = "00304593"
	adk.NAMA_PEGAWAI = "Gita Arifatun Nisa"
	adk.PDWK.limit = 100000000

	fmt.Println(adk.GetUser())
	fmt.Println(adk.GetPDWK())

	GetUserDetails(adk)
}
