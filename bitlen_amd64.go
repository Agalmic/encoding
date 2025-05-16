//go:build amd64 && !gccgo
// +build amd64,!gccgo

package encoding

import "math/bits"

// bitlen returns the number of bits required to represent x.
func bitlen(x uint64) int {
	return bits.Len64(x)
}
