// +build without_openssl

package stupidgcm

import (
	"os"

	"github.com/rfjakob/gocryptfs/internal/tlog"
)

type stupidGCM struct{}

const (
	// BuiltWithoutOpenssl indicates if openssl been disabled at compile-time
	BuiltWithoutOpenssl = true
)

func errExit() {
	tlog.Fatal.Println("gocryptfs has been compiled without openssl support but you are still trying to use openssl")
	os.Exit(2)
}

func New(_ []byte) stupidGCM {
	errExit()
	// Never reached
	return stupidGCM{}
}

func (g stupidGCM) NonceSize() int {
	errExit()
	return -1
}

func (g stupidGCM) Overhead() int {
	errExit()
	return -1
}

func (g stupidGCM) Seal(_, _, _, _ []byte) []byte {
	errExit()
	return nil
}

func (g stupidGCM) Open(_, _, _, _ []byte) ([]byte, error) {
	errExit()
	return nil, nil
}
