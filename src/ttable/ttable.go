package ttable

// Encrypts a 16-bytes block src using AES encryption with
// extended key xk and for nr rounds.
// It returns the encrypted block.
func Encrypt(src, key []byte, nr int) []byte {
	var tmp0, tmp1, tmp2, tmp3 uint32

	xk := KeyExpansion(key, nr)

	// Initialize the state, stored in s0, s1, s2 and s3
	s0 := uint32(src[0])<<24 | uint32(src[1])<<16 | uint32(src[2])<<8 | uint32(src[3])
	s1 := uint32(src[4])<<24 | uint32(src[5])<<16 | uint32(src[6])<<8 | uint32(src[7])
	s2 := uint32(src[8])<<24 | uint32(src[9])<<16 | uint32(src[10])<<8 | uint32(src[11])
	s3 := uint32(src[12])<<24 | uint32(src[13])<<16 | uint32(src[14])<<8 | uint32(src[15])

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
	dst[0], dst[1], dst[2], dst[3] = byte(s0>>24), byte(s0>>16), byte(s0>>8), byte(s0)
	dst[4], dst[5], dst[6], dst[7] = byte(s1>>24), byte(s1>>16), byte(s1>>8), byte(s1)
	dst[8], dst[9], dst[10], dst[11] = byte(s2>>24), byte(s2>>16), byte(s2>>8), byte(s2)
	dst[12], dst[13], dst[14], dst[15] = byte(s3>>24), byte(s3>>16), byte(s3>>8), byte(s3)

	return dst
}
