package ttable

import "encoding/binary"

// Encrypts a 16-bytes block src using AES encryption with
// extended key xk and for nr rounds.
// It returns the encrypted block.w
func Encrypt(src, key []byte, nr int) []byte {
	var tmp0, tmp1, tmp2, tmp3 uint32

	xk := KeyExpansion(key, nr)

	// Initialize the state, stored in s0, s1, s2 and s3
	s0 := binary.BigEndian.Uint32(src[0:4])
	s1 := binary.BigEndian.Uint32(src[4:8])
	s2 := binary.BigEndian.Uint32(src[8:12])
	s3 := binary.BigEndian.Uint32(src[12:16])

	// Add the first round key to the state
	s0 ^= xk[0]
	s1 ^= xk[1]
	s2 ^= xk[2]
	s3 ^= xk[3]

	for i := 1; i < nr; i++ {
		// This performs SubBytes + ShiftRows + MixColumns + AddRoundKey
		tmp0 = te0[s0>>24] ^ te1[s1>>16&0xff] ^ te2[s2>>8&0xff] ^ te3[s3&0xff] ^ xk[4*i]
		tmp1 = te0[s1>>24] ^ te1[s2>>16&0xff] ^ te2[s3>>8&0xff] ^ te3[s0&0xff] ^ xk[4*i+1]
		tmp2 = te0[s2>>24] ^ te1[s3>>16&0xff] ^ te2[s0>>8&0xff] ^ te3[s1&0xff] ^ xk[4*i+2]
		tmp3 = te0[s3>>24] ^ te1[s0>>16&0xff] ^ te2[s1>>8&0xff] ^ te3[s2&0xff] ^ xk[4*i+3]

		s0, s1, s2, s3 = tmp0, tmp1, tmp2, tmp3
	}

	// Perform the last SubBytes + ShiftRows + AddRoundKey
	tmp0 = (uint32(sbox[s0>>24])<<24 | uint32(sbox[s1>>16&0xff])<<16 | uint32(sbox[s2>>8&0xff])<<8 | uint32(sbox[s3&0xff])) ^ xk[4*nr]
	tmp1 = (uint32(sbox[s1>>24])<<24 | uint32(sbox[s2>>16&0xff])<<16 | uint32(sbox[s3>>8&0xff])<<8 | uint32(sbox[s0&0xff])) ^ xk[4*nr+1]
	tmp2 = (uint32(sbox[s2>>24])<<24 | uint32(sbox[s3>>16&0xff])<<16 | uint32(sbox[s0>>8&0xff])<<8 | uint32(sbox[s1&0xff])) ^ xk[4*nr+2]
	tmp3 = (uint32(sbox[s3>>24])<<24 | uint32(sbox[s0>>16&0xff])<<16 | uint32(sbox[s1>>8&0xff])<<8 | uint32(sbox[s2&0xff])) ^ xk[4*nr+3]
	s0, s1, s2, s3 = tmp0, tmp1, tmp2, tmp3

	// Unpack the state into the destination block. For performance purpose dst could be passed to the function
	// rather than being allocated each time.
	dst := make([]byte, 16)
	binary.BigEndian.PutUint32(dst[0:4], s0)
	binary.BigEndian.PutUint32(dst[4:8], s1)
	binary.BigEndian.PutUint32(dst[8:12], s2)
	binary.BigEndian.PutUint32(dst[12:16], s3)

	return dst
}
