package main

import "github.com/mr-tron/base58"

const (
	AddressPrefix  = "G"
	AddressLength  = 36
	CheckSumLength = 4
)

type Address [AddressLength]byte

func (a Address) Bytes() []byte {
	return a[:]
}

func (a *Address) SetBytes(in []byte) {
	if len(in) > len(a) {
		in = in[len(in)-AddressLength:]
	}
	copy(a[AddressLength-len(in):], in)
}

func BytesToAddress(b []byte) Address {
	addr := Address{}
	addr.SetBytes(b)
	return addr
}

func (a Address) String() string {
	return AddressPrefix + base58.Encode(a[:])
}

func (a Address) Raw() []byte {
	return a[:AddressLength-CheckSumLength]
}
