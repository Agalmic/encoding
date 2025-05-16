//go:build !gccgo && (386 || arm)
// +build !gccgo
// +build 386 arm

package encoding

// This is defined in util_{amd64,386}.s, copied from pkg/math/big/arith_{amd64/386}.s
func bitlen(x uint64) (n int)
