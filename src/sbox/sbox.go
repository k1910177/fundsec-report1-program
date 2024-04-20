package sbox

import "encoding/binary"

// Encrypts a 16-bytes block src using AES encryption with
// extended key xk and for nr rounds.
// It returns the encrypted block.
func Encrypt(src, key []byte, nr int) []byte {
	xk := KeyExpansion(key, nr)

	state := make([]uint32, 4)

	// Initialize the state with the plaintext block
	for i := 0; i < 4; i++ {
		state[i] = binary.BigEndian.Uint32(src[4*i:])
	}

	// Add the first round key to the state
	AddRoundKey(xk[0:4], state)

	for i := 1; i < nr; i++ {
		SubBytes(state)
		ShiftRows(state)
		MixColumns(state)
		AddRoundKey(xk[i*4:(i+1)*4], state)
	}

	SubBytes(state)
	ShiftRows(state)
	AddRoundKey(xk[nr*4:(nr+1)*4], state)

	dst := make([]byte, 16)
	for i := 0; i < 4; i++ {
		binary.BigEndian.PutUint32(dst[4*i:4*(i+1)], state[i])
	}

	return dst
}
